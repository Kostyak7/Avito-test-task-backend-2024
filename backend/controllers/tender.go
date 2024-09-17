package controllers

import (
	"backend/models"
	"backend/services"

	"encoding/json"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTender(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var tender models.Tender
        if err := json.NewDecoder(c.Request.Body).Decode(&tender); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
            return
        }
    
        newTender, err := services.CreateTenderInDB(db, tender)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error creating tender"})
            return
        }
    
        response, err := json.Marshal(newTender)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
            return
        }
        c.JSON(http.StatusOK, response)
    }
}

func GetTenders(db *gorm.DB) gin.HandlerFunc {
    return func (c *gin.Context) {
    limit := c.Request.URL.Query().Get("limit")
    offset := c.Request.URL.Query().Get("offset")
    serviceType := c.Request.URL.Query().Get("service_type")

    if !(serviceType == "Construction" || serviceType == "Delivery" || serviceType == "Manufacture") {
        c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid service type"})
        return
    }

    tenders, err := services.GetTendersFromDB(db, limit, offset, serviceType)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error fetching tenders"})
        return
    }

    response, err := json.Marshal(tenders)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
        return
    }
    c.JSON(http.StatusOK, response)
    }
}

func GetUserTenders(db *gorm.DB) gin.HandlerFunc {
return func (c *gin.Context) {
    username := c.Request.URL.Query().Get("username")
    limit := c.Request.URL.Query().Get("limit")
    offset := c.Request.URL.Query().Get("offset")

    tenders, err := services.GetUserTendersFromDB(db, username, limit, offset)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error fetching user tenders"})
        return
    }

    response, err := json.Marshal(tenders)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
        return
    }
    c.JSON(http.StatusOK, response)
}
}

func GetTenderStatus(db *gorm.DB) gin.HandlerFunc {
    return func (c *gin.Context) {
        tenderId := c.Param("tenderId")
        username := c.Request.URL.Query().Get("username")
    
        status, err := services.GetTenderStatusFromDB(db, tenderId, username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error fetching tender status"})
            return
        }
    
        response, err := json.Marshal(status)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
            return
        }
        c.JSON(http.StatusOK, response)
}
}

func UpdateTenderStatus(db *gorm.DB) gin.HandlerFunc {
    return func (c *gin.Context) {
        tenderId := c.Param("tenderId")
        status := c.Request.URL.Query().Get("status")
        username := c.Request.URL.Query().Get("username")

        if !(status == "Created" || status == "Published" || status == "Closed") {
            c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid service type"})
            return
        }
    
        updatedTender, err := services.UpdateTenderStatusInDB(db, tenderId, status, username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error updating tender status"})
            return
        }
    
        response, err := json.Marshal(updatedTender)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
            return
        }
        c.JSON(http.StatusOK, response)
}
}

func EditTender(db *gorm.DB) gin.HandlerFunc {
return func(c *gin.Context) {
    tenderId := c.Param("tenderId")
    username := c.Request.URL.Query().Get("username")

    var updates models.Tender
    if err := json.NewDecoder(c.Request.Body).Decode(&updates); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request body"})
        return
    }

    updatedTender, err := services.EditTenderInDB(db, tenderId, username, updates)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error editing tender"})
        return
    }

    response, err := json.Marshal(updatedTender)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
        return
    }
    c.JSON(http.StatusOK, response)
}
}

func RollbackTender(db *gorm.DB) gin.HandlerFunc {
return func(c *gin.Context) {
    tenderId := c.Param("tenderId") 
    version := c.Param("version")
    username := c.Request.URL.Query().Get("username")

    rolledBackTender, err := services.RollbackTenderInDB(db, tenderId, version, username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error rolling back tender"})
        return
    }

    response, err := json.Marshal(rolledBackTender)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"Error": "Error marshaling response"})
        return
    }
    c.JSON(http.StatusOK, response)
}
}
