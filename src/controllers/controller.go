package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) GetIndex(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Gin application!", "timestamp": time.Now().Unix()})
}

func (c *Controller) PostData(ctx *gin.Context) {
	var jsonData map[string]interface{}
	if err := ctx.ShouldBindJSON(&jsonData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"received": jsonData})
}

func (c *Controller) HealthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}
