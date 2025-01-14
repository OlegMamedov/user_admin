package CRUD

import (
	"first_app/src/database"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context){
	var user_id = c.DefaultQuery("user_id", "")

	if user_id == "" {
		c.JSON(400, gin.H{"message": "User ID is required"})
		return
	}

 	var user database.Users

	if err := database.DB.Where("id = ?", user_id).Delete(&user).Error; err != nil{
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	c.JSON(200, gin.H{"status":"User success deleted"})

}