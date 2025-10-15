package controller

import (
	"github.com/gin-gonic/gin"
)

func FindUserById(c *gin.Context) {
	userId := c.Param("userId")
	// Logic to find user by ID
	c.JSON(200, gin.H{"userId": userId})
}

func FindUserByEmail(c *gin.Context) {
	userEmail := c.Param("userEmail")
	// Logic to find user by email
	c.JSON(200, gin.H{"userEmail": userEmail})
}
