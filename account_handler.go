package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterAccountRoutes(r *gin.Engine) {
	r.POST("/api/create_account", handleCreateAccount)
}

func handleCreateAccount(c *gin.Context) {
	var createAccountReq CreateAccountRequest
	if err := c.ShouldBindJSON(&createAccountReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if GetAccountByEmail(createAccountReq.Email) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	account := CreateAccount(createAccountReq.Email, createAccountReq.Password)
	c.JSON(http.StatusOK, gin.H{"message": "Account created", "account_id": account.ID})
}
