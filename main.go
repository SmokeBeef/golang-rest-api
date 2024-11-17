package main

import (
	"dashboardapi/app"
	"dashboardapi/config"
	"dashboardapi/db"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	config.Run()
	if err != nil {
		panic("Error loading.env file")
	}
	db.RunDb()
	r := app.Routes()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Welcome to the API",
		})
	})
	r.Run(":8080")
}
