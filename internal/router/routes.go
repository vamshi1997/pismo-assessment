package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vamshi1997/pismo-assessment/internal/controller"
	"github.com/vamshi1997/pismo-assessment/internal/repo"
	"gorm.io/gorm"
)

func InitAppRoutes(router *gin.Engine, db *gorm.DB) {

	newRepo := repo.NewRepository(db)

	newController := controller.NewController(newRepo)

	router.GET("/status", controller.Status)
	router.POST("/accounts", newController.CreateAccount)
	router.GET("/accounts/:accountId", newController.GetAccount)
	router.POST("/transactions", newController.CreateTransaction)
}
