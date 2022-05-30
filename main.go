package main

import (
	"assignment3/config"
	"assignment3/controllers"
	"assignment3/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	webService := services.NewWebService()
	webController := controllers.NewWebController(webService)

	router.GET("/status", webController.GetStatus)
	router.Static("/pictures", "./pictures")
	router.Static("/js", "./js")

	router.Run(config.APP_PORT)
}
