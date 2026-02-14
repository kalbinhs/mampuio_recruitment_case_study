package handlers

import (
	"net/http"
	"wallet-api/db"
	"wallet-api/models"

	"github.com/gin-gonic/gin"
)

func GetBalance(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"response_code":    http.StatusNotFound,
			"response_message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":          user.ID,
		"user_name":        user.Name,
		"balance":          user.Balance,
		"response_code":    http.StatusOK,
		"response_message": "Balance retrieved",
	})
}
