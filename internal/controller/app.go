package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"github.com/vamshi1997/pismo-assessment/internal/repo"
	"net/http"
	"strconv"
)

func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

func CreateAccount(ctx *gin.Context) {
	var account model.Account

	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	repo.CreateAccount(account)

	ctx.JSON(http.StatusOK, gin.H{"document_number": account.DocumentNumber})
}

func GetAccount(ctx *gin.Context) {
	accountID, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account ID"})
		return
	}

	accountInfo, err := repo.GetAccount(accountID)
	if err != nil || accountInfo == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"account_id": accountInfo.ID, "document_number": accountInfo.DocumentNumber})
}

func CreateTransaction(ctx *gin.Context) {
	var transaction model.Transaction

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := repo.CreateTransaction(transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction"})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}
