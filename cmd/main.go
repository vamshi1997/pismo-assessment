package main

import (
	"github.com/vamshi1997/pismo-assessment/internal/boot"
	"github.com/vamshi1997/pismo-assessment/internal/router"
	"log"
)

func main() {
	log.Println("Starting Go Web Application")
	boot.InitApp()
	router.InitiateRouter(boot.GetDB())
}
