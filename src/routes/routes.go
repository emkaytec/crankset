package routes

import (
	"emkaytec.io/crankset/src/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	controller := controllers.Controller{}

	router.GET("/", controller.GetIndex)
	router.POST("/data", controller.PostData)

	router.GET("/ip", controller.GetIPAddress)
	router.GET("/healthz", controller.HealthCheck)
	router.GET("/health_check", controller.HealthCheck)
}
