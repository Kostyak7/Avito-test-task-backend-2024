package services

import (
    "backend/models"
    
    "strconv"
    "errors"
    "fmt"
    "time"
    "github.com/google/uuid"
	"gorm.io/gorm"	
)

func CreateBid(db *gorm.DB, bid models.Bid) (*models.Bid, error) {
    if bid.Name == "" || bid.Description == "" || bid.TenderID == uuid.Nil || bid.AuthorType == "" || bid.AuthorID == uuid.Nil {
        return nil, errors.New("missing required fields")
    }

    createdBid := bid
    createdBid.ID = uuid.New()
    createdBid.CreatedAt = time.Now()

    err := db.Create(&createdBid).Error
    if err != nil {
        return nil, err
    }

    return &createdBid, nil
}

func GetUserBids(db *gorm.DB, username, limit_str, offset_str string) ([]models.Bid, error) {
    var limit int = 1000
    var offset int
    var errl, erro error
    if limit_str != "" {
    limit, errl = strconv.Atoi(limit_str)
    if errl != nil {
        return nil, errl
    }
    }
    if offset_str != "" {
    offset, erro = strconv.Atoi(offset_str)
    if erro != nil {
        return nil, erro
    }
    }

    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bids []models.Bid
    err := db.Where("author_id = ?", user.ID).
        Limit(limit). 
        Offset(offset). 
        Order("created_at ASC").
        Find(&bids).Error

    if err != nil {
        return nil, err
    }

    return bids, nil
}

func GetBidsForTender(db *gorm.DB, tenderID string, username string, limit_str string, offset_str string) ([]models.Bid, error) {
    var limit int = 1000
    var offset int
    var errl, erro error
    if limit_str != "" {
    limit, errl = strconv.Atoi(limit_str)
    if errl != nil {
        return nil, errl
    }
    }
    if offset_str != "" {
    offset, erro = strconv.Atoi(offset_str)
    if erro != nil {
        return nil, erro
    }
    }

    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bids []models.Bid
    err := db.Where("tender_id = ?", tenderID).
        Where("author_id = ?", user.ID).
        Limit(limit).
        Offset(offset).
        Order("created_at ASC").
        Find(&bids).Error

    if err != nil {
        return nil, err
    }

    return bids, nil
}

func GetBidStatus(db *gorm.DB, bidID string, username string) (string, error) {
    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return "", fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bid models.Bid
    err := db.Where("id = ? AND author_id = ?", bidID, user.ID).First(&bid).Error

    if err != nil {
        return "", err
    }

    return bid.Status, nil
}

func UpdateBidStatus(db *gorm.DB, bidID string, status string, username string) (*models.Bid, error) {
    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bid models.Bid
    err := db.Where("id = ? AND author_id = ?", bidID, user.ID).First(&bid).Error

    if err != nil {
        return nil, err
    }

    bid.Status = status
    err = db.Save(&bid).Error

    if err != nil {
        return nil, err
    }

    return &bid, nil
}

func EditBid(db *gorm.DB, bidID string, username string, updates models.Bid) (*models.Bid, error) {
    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bid models.Bid
    err := db.Where("id = ? AND author_id = ?", bidID, user.ID).First(&bid).Error

    if err != nil {
        return nil, err
    }

    if updates.Name != "" {
        bid.Name = updates.Name
    }
    if updates.Description != "" {
        bid.Description = updates.Description
    }

    err = db.Save(&bid).Error

    if err != nil {
        return nil, err
    }

    return &bid, nil
}

func SubmitBidDecision(db *gorm.DB, bidID string, decision string, username string) (*models.Bid, error) {
    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bid models.Bid
    err := db.Where("id = ? AND author_id = ?", bidID, user.ID).First(&bid).Error

    if err != nil {
        return nil, err
    }

    bid.Decision = decision
    err = db.Save(&bid).Error

    if err != nil {
        return nil, err
    }

    return &bid, nil
}

func SubmitBidFeedback(db *gorm.DB, bidID string, feedback string, username string) (*models.Bid, error) {
    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bid models.Bid
    err := db.Where("id = ? AND author_id = ?", bidID, user.ID).First(&bid).Error

    if err != nil {
        return nil, err
    }

    bid.Feedback = feedback
    err = db.Save(&bid).Error

    if err != nil {
        return nil, err
    }

    return &bid, nil
}

func RollbackBid(db *gorm.DB, bidID string, version string, username string) (*models.Bid, error) {
    var user models.Employee
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", username)
    }

    var bid models.Bid
    err := db.Where("id = ? AND author_id = ?", bidID, user.ID).First(&bid).Error

    if err != nil {
        return nil, err
    }

    err = db.Where("bid_id = ? AND version = ?", bidID, version).First(&bid).Error
    if err != nil {
        return nil, err
    }

    bid.Version++
    err = db.Save(&bid).Error

    if err != nil {
        return nil, err
    }

    return &bid, nil
}

func GetBidReviews(db *gorm.DB, tenderID string, authorUsername string, requesterUsername string, limit_str string, offset_str string) ([]models.BidReview, error) {
    var limit int = 1000
    var offset int
    var errl, erro error
    if limit_str != "" {
    limit, errl = strconv.Atoi(limit_str)
    if errl != nil {
        return nil, errl
    }
    }
    if offset_str != "" {
    offset, erro = strconv.Atoi(offset_str)
    if erro != nil {
        return nil, erro
    }
    }

    var authir_user models.Employee
    if err := db.Where("username = ?", authorUsername).First(&authir_user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", authorUsername)
    }

    var request_user models.Employee
    if err := db.Where("username = ?", requesterUsername).First(&request_user).Error; err != nil {
        return nil, fmt.Errorf("user with username '%s' does not exist", requesterUsername)
    }
    
    var reviews []models.BidReview
    err := db.Joins("JOIN bids ON bids.id = bid_reviews.bid_id").
        Where("bids.tender_id = ? AND bids.author_id = ?", tenderID, authir_user.ID).
        Limit(limit).
        Offset(offset).
        Find(&reviews).Error
    
    if err != nil {
        return nil, fmt.Errorf("error fetching reviews: %v", err)
    }

    return reviews, nil
}