package routes

import (
	"first_app/src/database"

	"github.com/gin-gonic/gin"
)

func GetListUsers(c *gin.Context) {
	var users []database.Users

	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(404, gin.H{"error": "not found"})
		return
	}

	c.JSON(200, users)
}