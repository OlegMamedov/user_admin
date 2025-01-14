package database

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name  string `gorm:"size:255;not null"`
	Password string `gorm:"size:255;not null"`
    Email string `gorm:"size:255;unique;not null"`
    Age   int    `gorm:"not null"`
}