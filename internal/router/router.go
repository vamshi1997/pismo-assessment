package router

import (
	"fmt"
	"log"

	"github.com/vamshi1997/pismo-assessment/internal/boot"

	"github.com/gin-gonic/gin"
)

func InitiateRouter() {
	router := gin.New()

	cfg := boot.GetConfig()

	InitAppRoutes(router)

	serverAddr := fmt.Sprintf("%s:%v", cfg.AppConfig.Server.Host, cfg.AppConfig.Server.Port)
	err := router.Run(serverAddr)
	if err != nil {
		log.Println("error while running the server:", err)
		return
	}

	log.Printf("Server Started Successfully & listening to Port: %v", cfg.AppConfig.Server.Port)
}
