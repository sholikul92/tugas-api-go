package main

import (
	"project-api-go/controllers/routes"
	"project-api-go/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDB()

	app := gin.Default()

	// endpoint
	app.POST("/users", routes.AddUser)
	app.GET("/users", routes.GetAllUsers)
	app.GET("/users/:id", routes.GetUserById)
	app.PUT("/users/:id", routes.UpdateUser)
	app.DELETE("users/:id", routes.DeleteUser)

	app.Run(":9000")
}
