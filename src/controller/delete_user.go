package controller

import "github.com/gin-gonic/gin"

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	// Logic to delete user by ID
	c.JSON(200, gin.H{"message": "User deleted successfully", "userId": userId})
}
