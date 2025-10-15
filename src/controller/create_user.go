package controller

import "github.com/gin-gonic/gin"

func CreateUser(c *gin.Context) {
	// Logic to create a new user
	c.JSON(201, gin.H{"message": "User created successfully"})
}
