package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kunalvirwal/Velocity-Load-Balancer/internal/state"
)

func GetAllServices(c *gin.Context) {
	fmt.Println(state.AllServers)
}

func Handle404(c *gin.Context) {
	c.JSON(404, gin.H{"status": "error", "message": "not found"})
}
