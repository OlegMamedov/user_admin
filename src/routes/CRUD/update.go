package CRUD

import (
	"first_app/src/database"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var user_id = c.DefaultQuery("user_id", "")

	var user UserUpdate

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	if err := database.DB.Model(database.Users{}).Where("id = ?", user_id).Updates(map[string]interface{}{
		"name": user.Name,
		"age":  user.Age,
		"email": user.Email,
		"password": user.Password,
	}).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "User updated",
		"user": user,
	})
}
