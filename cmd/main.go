package main

import (
	"github.com/vamshi1997/pismo-assessment/internal/boot"
	"github.com/vamshi1997/pismo-assessment/internal/router"
)

func main() {
	boot.InitApp()
	router.InitiateRouter()
}
