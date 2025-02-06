package router

import (
	"fmt"
	"gorm.io/gorm"
	"log"

	"github.com/vamshi1997/pismo-assessment/internal/boot"

	"github.com/gin-gonic/gin"
)

func InitiateRouter(db *gorm.DB) {
	router := gin.New()

	cfg := boot.GetConfig()

	InitAppRoutes(router, db)

	serverAddr := fmt.Sprintf("%s:%v", cfg.AppConfig.Server.Host, cfg.AppConfig.Server.Port)
	err := router.Run(serverAddr)
	if err != nil {
		log.Println("error while running the server:", err)
		return
	}

	log.Printf("Server Started Successfully & listening to Port: %v", cfg.AppConfig.Server.Port)
}
