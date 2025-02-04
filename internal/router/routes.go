package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vamshi1997/pismo-assessment/internal/controller"
)

func InitAppRoutes(router *gin.Engine) {
	router.GET("/status", controller.Status)
	router.POST("/accounts", controller.CreateAccount)
	router.GET("/accounts/:accountId", controller.GetAccount)
	router.POST("/transactions", controller.CreateTransaction)
}
