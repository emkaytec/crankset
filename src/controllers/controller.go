package controllers

import (
	"net"
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

func (c *Controller) GetIPAddress(ctx *gin.Context) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get IP address"})
		return
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ctx.JSON(http.StatusOK, gin.H{"ip_address": ipNet.IP.String()})
				return
			}
		}
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No IP address found"})
}
