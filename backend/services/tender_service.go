package services

import (
    "backend/models"
    
    "fmt"
    "gorm.io/gorm"
    "strconv"
)


func GetTendersFromDB(db *gorm.DB, limit_str, offset_str string, serviceType string) ([]models.Tender, error) {
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

    var tenders []models.Tender
    if err := db.Where("service_type == ?", serviceType).Limit(limit).Offset(offset).Order("name ASC").Find(&tenders).Error; err != nil {
        return nil, err
    }
    return tenders, nil
}

func CreateTenderInDB(db *gorm.DB, tender models.Tender) (models.Tender, error) {
    var existingTender models.Tender
    if err := db.Where("name = ?", tender.Name).First(&existingTender).Error; err == nil {
        return tender, fmt.Errorf("tender with name '%s' already exists", tender.Name)
    }

    var user models.Employee
    if err := db.Where("username = ?", tender.CreatorUsername).First(&user).Error; err != nil {
        return tender, fmt.Errorf("user with username '%s' does not exist", tender.CreatorUsername)
    }

    var organization models.Organization
    if err := db.Where("id = ?", tender.OrganizationID).First(&organization).Error; err != nil {
        return tender, fmt.Errorf("organization with id '%d' does not exist", tender.OrganizationID)
    }

    if err := db.Create(&tender).Error; err != nil {
        return tender, err
    }
    return tender, nil
}

func GetUserTendersFromDB(db *gorm.DB, username string, limit_str, offset_str string) ([]models.Tender, error) {
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

    var tenders []models.Tender
    if err := db.Where("creator_username = ?", username).
        Limit(limit).
        Offset(offset).
        Order("name ASC").
        Find(&tenders).Error; err != nil {
        return nil, err
    }
    return tenders, nil
}

func GetTenderStatusFromDB(db *gorm.DB, tenderId, username string) (string, error) {
    var status string
    if err := db.Model(&models.Tender{}).
        Where("id = ? AND creator_username = ?", tenderId, username).
        Select("status").
        Scan(&status).Error; err != nil {
        return "", err
    }
    return status, nil
}

func UpdateTenderStatusInDB(db *gorm.DB, tenderId, status, username string) (models.Tender, error) {
    var tender models.Tender
    if err := db.Model(&tender).
        Where("id = ? AND creator_username = ?", tenderId, username).
        Updates(models.Tender{Status: status}).Error; err != nil {
        return tender, err
    }

    if err := db.Where("id = ?", tenderId).First(&tender).Error; err != nil {
        return tender, err
    }
    return tender, nil
}

func EditTenderInDB(db *gorm.DB, tenderId, username string, updates models.Tender) (models.Tender, error) {
    var tender models.Tender

    if err := db.Model(&tender).
        Where("id = ? AND creator_username = ?", tenderId, username).
        Updates(updates).Error; err != nil {
        return tender, err
    }

    if err := db.Where("id = ?", tenderId).First(&tender).Error; err != nil {
        return tender, err
    }

    return tender, nil
}

func RollbackTenderInDB(db *gorm.DB, tenderId, version_str, username string) (models.Tender, error) {
    var previousVersion models.Tender

    version, errv := strconv.Atoi(version_str)
    if errv != nil {        
        return previousVersion, errv
    }

    if err := db.Where("tender_id = ? AND version = ? AND creator_username = ?", tenderId, version, username).
        First(&previousVersion).Error; err != nil {
        return previousVersion, err
    }

    if err := db.Model(&models.Tender{}).
        Where("id = ? AND creator_username = ?", tenderId, username).
        Updates(models.Tender{Name: previousVersion.Name, Description: previousVersion.Description, ServiceType: previousVersion.ServiceType}).
        Error; err != nil {
        return previousVersion, err
    }

    var rolledBackTender models.Tender
    if err := db.Where("id = ?", tenderId).First(&rolledBackTender).Error; err != nil {
        return rolledBackTender, err
    }

    return rolledBackTender, nil
}