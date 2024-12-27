package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Controller struct{}

func (c *Controller) GetIndex(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the Gin application!"})
}

func (c *Controller) PostData(ctx *gin.Context) {
    var jsonData map[string]interface{}
    if err := ctx.ShouldBindJSON(&jsonData); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"received": jsonData})
}