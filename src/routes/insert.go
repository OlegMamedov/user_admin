package routes

import (
	"first_app/src/database"

	"github.com/gin-gonic/gin"
)


func InsertUser(c *gin.Context) {
    var user database.Users

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    if err := database.DB.Create(&user).Error; err != nil {
        c.JSON(500, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(200, gin.H{"message": "User created successfully", "user": user})
}