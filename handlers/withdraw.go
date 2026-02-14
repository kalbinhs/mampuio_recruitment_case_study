package handlers

import (
	"net/http"
	"wallet-api/db"
	"wallet-api/models"

	"github.com/gin-gonic/gin"
)

type WithdrawRequest struct {
	UserID uint  `json:"user_id"`
	Amount int64 `json:"amount"`
}

func Withdraw(c *gin.Context) {
	var req WithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code":    http.StatusBadRequest,
			"response_message": "Invalid request",
		})
		return
	}

	var user models.User
	if err := db.DB.First(&user, req.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"response_code":    http.StatusNotFound,
			"response_message": "User not found",
		})
		return
	}

	if req.Amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code":    http.StatusBadRequest,
			"response_message": "Amount must be positive",
		})
		return
	}

	if user.Balance < req.Amount {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code":    http.StatusBadRequest,
			"response_message": "Insufficient balance",
		})
		return
	}

	user.Balance -= req.Amount
	db.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"user_id":          user.ID,
		"user_name":        user.Name,
		"balance":          user.Balance,
		"response_code":    http.StatusOK,
		"response_message": "Withdraw successful",
	})
}
