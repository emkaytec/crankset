package routes

import (
    "github.com/gin-gonic/gin"
    "my-gin-app/src/controllers"
)

func SetupRoutes(router *gin.Engine) {
    controller := controllers.Controller{}

    router.GET("/", controller.GetIndex)
    router.POST("/data", controller.PostData)
}