package main

import (
	"log"
	"wallet-api/db"
	"wallet-api/handlers"
	"wallet-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatal("Failed to connect database")
	}

	db.DB.AutoMigrate(&models.User{})

	r := gin.Default()
	r.GET("/balance/:id", handlers.GetBalance)
	r.POST("/withdraw", handlers.Withdraw)

	r.Run(":8080")
}
