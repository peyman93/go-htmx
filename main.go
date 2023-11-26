package main

import (
	"github.com/gin-gonic/gin"
	"github.com/peyman93/go-jwt/controllers"
	"github.com/peyman93/go-jwt/initializers"
	"log"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {

	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	err := r.Run()
	if err != nil {
		log.Fatal("Server not started")
	}
}
