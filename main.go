package main

import (
	"first_app/src/database"
	"first_app/src/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("Server start on: http://localhost:8080")

	database.ConnectDatabase()

	r.POST("/register", routes.InsertUser)
	r.GET("/list", routes.GetListUsers)
	r.DELETE("/delete", routes.DeleteUser)
	r.PUT("/update", routes.UpdateUser)

	r.Run()


}

