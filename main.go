package main

import (
	"first_app/src/database"
	"first_app/src/routes/CRUD"
	"first_app/src/routes/auth"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("Server start on: http://localhost:8080")

	database.ConnectDatabase()

	r.POST("/create", CRUD.InsertUser)
	r.GET("/list", CRUD.GetListUsers)
	r.DELETE("/delete", CRUD.DeleteUser)
	r.PUT("/update", CRUD.UpdateUser)
	r.POST("/login", auth.Auth)
	r.POST("/register", auth.Register)

	r.Run()


}

