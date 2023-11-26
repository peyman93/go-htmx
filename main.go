package main

import (
	"github.com/gin-gonic/gin"
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

	r.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.Run()
	if err != nil {
		log.Fatal("Server not started")
	}
}
