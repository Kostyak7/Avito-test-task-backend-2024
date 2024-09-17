package controllers

import (
    "net/http"
    "backend/models"
    "backend/services"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "encoding/json"
    "github.com/google/uuid"
)


func CreateBID(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
      var bid models.Bid

    if err := json.NewDecoder(c.Request.Body).Decode(&bid); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
        return
    }

    if bid.Name == "" || bid.Description == "" || bid.TenderID == uuid.Nil || bid.AuthorType == "" || bid.AuthorID == uuid.Nil {
      c.JSON(http.StatusBadRequest, gin.H{"Error": "Missing required fields"})
        return
    }

    createdBid, err := services.CreateBid(db, bid)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, createdBid)
    }
}

func GetUserBids(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    username := c.Request.URL.Query().Get("username")
    limit := c.Request.URL.Query().Get("limit")
    offset := c.Request.URL.Query().Get("offset")



    bids, err := services.GetUserBids(db, username, limit, offset)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, bids)
}
}

func GetTenderBids(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    tenderID := c.Param("id")
    username := c.Request.URL.Query().Get("username")
    limit := c.Request.URL.Query().Get("limit")
    offset := c.Request.URL.Query().Get("offset")

    bids, err := services.GetBidsForTender(db, tenderID, username, limit, offset)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, bids)
}
}

func EditBID(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    bidID := c.Param("id")
    username := c.Request.URL.Query().Get("username")

    var updates models.Bid
    if err := json.NewDecoder(c.Request.Body).Decode(&updates); err != nil {
      c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
        return
    }

    updatedBid, err := services.EditBid(db, bidID, username, updates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedBid)
}
}

func RollbackBID(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    bidID := c.Param("id")
    version := c.Param("version")
    username := c.Request.URL.Query().Get("username")

    updatedBid, err := services.RollbackBid(db, bidID, version, username)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, updatedBid)
}
}

func GetBIDReviews(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
    tenderID := c.Param("id")
    authorUsername := c.Request.URL.Query().Get("authorUsername")
    requesterUsername := c.Request.URL.Query().Get("requesterUsername")
    limit := c.Request.URL.Query().Get("limit")
    offset := c.Request.URL.Query().Get("offset")

    reviews, err := services.GetBidReviews(db, tenderID, authorUsername, requesterUsername, limit, offset)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, reviews)
}
}

func GetBidStatus(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
  bidID := c.Param("id")
  username := c.Request.URL.Query().Get("username")

  status, err := services.GetBidStatus(db, bidID, username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, status)
}
}

func UpdateBidStatus(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
  bidID := c.Param("id")
  status := c.Request.URL.Query().Get("status")
  username := c.Request.URL.Query().Get("username")

  if !(status == "Created" || status == "Published" || status == "Canceled") {
    c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid service type"})
    return
}

  updatedBid, err := services.UpdateBidStatus(db, bidID, status, username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, updatedBid)
}
}

func SubmitBidDecision(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
  bidID := c.Param("id")
  decision := c.Request.URL.Query().Get("decision")
  username := c.Request.URL.Query().Get("username")

  updatedBid, err := services.SubmitBidDecision(db, bidID, decision, username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, updatedBid)
}
}

func SubmitBidFeedback(db *gorm.DB) gin.HandlerFunc {
  return func(c *gin.Context) {
  bidID := c.Param("id")
  feedback := c.Request.URL.Query().Get("feedback")
  username := c.Request.URL.Query().Get("username")

  updatedBid, err := services.SubmitBidFeedback(db, bidID, feedback, username)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
      return
  }

  c.JSON(http.StatusOK, updatedBid)
}
}
