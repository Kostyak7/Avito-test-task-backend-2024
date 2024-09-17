package routes

import (
    "backend/controllers"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
    router := gin.Default()

    api := router.Group("/api")
    {
        api.GET("/ping", controllers.Ping)

        // Tenders
        tenders := api.Group("/tenders")
        {
            tenders.GET("", controllers.GetTenders(db))
            tenders.POST("/new", controllers.CreateTender(db))
            tenders.GET("/my", controllers.GetUserTenders(db))

            tenders.GET("/:tenderId/status", controllers.GetTenderStatus(db))
            tenders.PUT("/:tenderId/status", controllers.UpdateTenderStatus(db))
            tenders.PATCH("/:tenderId/edit", controllers.EditTender(db))
            tenders.PUT("/:tenderId/rollback/:version", controllers.EditTender(db))
        }

        // BIDs
        bids := api.Group("/bids")
        {
            bids.POST("/new", controllers.CreateBID(db))
            bids.GET("/my", controllers.GetUserBids(db))

            bids.GET("/:id/status", controllers.GetBidStatus(db))
            bids.PUT("/:id/status", controllers.UpdateBidStatus(db))
            bids.PATCH("/:id/edit", controllers.EditBID(db))
            bids.PUT("/:id/submit_decision", controllers.SubmitBidDecision(db))
            bids.PUT("/:id/feedback", controllers.SubmitBidFeedback(db))
            bids.PUT("/:id/rollback/:version", controllers.RollbackBID(db))
            
            bids.GET("/:id/list", controllers.GetTenderBids(db))
            bids.GET("/:id/reviews", controllers.GetBIDReviews(db))
        }
    }

    return router
}
