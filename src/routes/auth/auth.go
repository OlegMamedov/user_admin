package auth

import (
	"errors"
	"first_app/src/database"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var user database.Users

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid required"})
		return
	}

	result := database.DB.Model(&database.Users{}).Where("email = ?", user.Email).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// Пользователь не найден, можно продолжать регистрацию
			log.Println("User not found, proceeding with registration")
		} else {
			// Другая ошибка, нужно обработать
			c.JSON(500, gin.H{"error": result.Error.Error()})
			return
		}
	} else {
		// Пользователь уже существует
		c.JSON(400, gin.H{"error": "User already exists"})
		return
	}

	HashedPass, err := HashPassword(user.Password)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Password = HashedPass

	if err := database.DB.Model(&database.Users{}).Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User register successfuly", "user_email": user.Email})
}

func Auth(c *gin.Context) {
	var user AuthSchema

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid required"})
		return
	}

	var dbUser database.Users

	if err := database.DB.Model(&database.Users{}).Where("email = ?", user.Email).First(&dbUser).Error; err != nil {
		c.JSON(400, gin.H{"error": "Invalid email"})
		return
	}

	if !CheckPasswordHash(user.Password, dbUser.Password){
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	token := GenerateToken()

	newToken := database.Tokens{
		Token: token,
		UserID: int(dbUser.ID),
	}

	if err := database.DB.Create(&newToken).Error; err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"access_token": token})


}