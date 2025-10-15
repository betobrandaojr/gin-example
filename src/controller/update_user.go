package controller

import "github.com/gin-gonic/gin"

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	// Logic to update user by ID
	c.JSON(200, gin.H{"message": "User updated successfully", "userId": userId})
}
