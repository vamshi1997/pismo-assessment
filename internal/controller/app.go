package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"github.com/vamshi1997/pismo-assessment/internal/repo"
	"net/http"
	"strconv"
)

type Controller struct {
	repo repo.IRepository
}

func NewController(repo repo.IRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

// Status method Gives application status to check if it's working or not
func Status(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]interface{}{"status": "ok"})
}

// CreateAccount method takes document number and create account accordingly
func (c *Controller) CreateAccount(ctx *gin.Context) {
	var (
		account     model.Account
		accountInfo model.Account
	)

	// decoding the request payload to account model
	if err := ctx.ShouldBindJSON(&account); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"error_msg": "invalid request data",
			"msg":       "Not able to create account",
		})
		return
	}

	// check if document number is valid or not
	if len(account.DocumentNumber) != 11 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_msg": "Document number given is not valid",
			"msg":       "Not able to create account",
		})
		return
	}

	accountInfo, err := c.repo.CreateAccount(account)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"error_msg": "Internal Server Error",
			"msg":       "Not able to create account",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"document_number": accountInfo.DocumentNumber,
		"account_id":      accountInfo.ID,
		"msg":             "Account created successfully",
	})
}

// GetAccount method takes account number and fetch account details if exists else return error
func (c *Controller) GetAccount(ctx *gin.Context) {

	// fetch account id from request param and converting it to integer
	accountID, err := strconv.Atoi(ctx.Param("accountId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"error_msg": "Not valid accountId",
			"msg":       "Not able to fetch account details",
		})
		return
	}

	// fetch account info from db
	accountInfo, err := c.repo.GetAccount(uint(accountID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":     err.Error(),
			"error_msg": "Account not found",
			"msg":       "Not able to fetch account details",
		})
		return
	}

	if &accountInfo == nil || accountInfo.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error_msg": "Account Details are empty",
			"msg":       "Not able to fetch account details",
		})
		return // Add this return statement
	}

	ctx.JSON(http.StatusOK, gin.H{
		"account_id":      accountInfo.ID,
		"document_number": accountInfo.DocumentNumber,
		"msg":             "Account details fetched successfully",
	})
}

// CreateTransaction method takes account_id, operation_type and amount and create it accordingly
func (c *Controller) CreateTransaction(ctx *gin.Context) {
	var (
		transaction     model.Transaction
		transactionInfo *model.Transaction
	)

	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error":     err.Error(),
			"error_msg": "Invalid request body",
			"msg":       "Not able to create transaction",
		})
		return
	}

	// check1: operation should be valid type
	if !model.IsValidOperationType(transaction.OperationTypeId) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_msg": "Invalid operation type",
			"msg":       "Not able to create transaction"})
		return
	}

	// check 2: purchase operations should have negative amount
	if (transaction.OperationTypeId == 1 || transaction.OperationTypeId == 2 || transaction.OperationTypeId == 3) && transaction.Amount >= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_msg": "Amount can not be positive for this operation type",
			"msg":       "Not able to create transaction",
		})
		return
	}

	// check 3: credit voucher should have positive amount
	if transaction.OperationTypeId == 4 && transaction.Amount < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error_mgs": "Amount can not be negative for this operation type",
			"msg":       "Not able to create transaction",
		})
		return
	}

	// check 4: if account is valid or not, then only transaction can be done
	accountInfo, err := c.repo.GetAccount(transaction.AccountID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":     err.Error(),
			"error_msg": "Account not found",
			"msg":       "Not able to create transaction",
		})
		return
	}
	if accountInfo == nil || accountInfo.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error_msg": "Account Details are empty",
			"msg":       "Not able to create transaction",
		})
		return // Add this return statement
	}

	if transactionInfo, err = c.repo.CreateTransaction(transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":     err.Error(),
			"error_msg": "Invalid transaction",
			"msg":       "Not able to create transaction"})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"msg":               "transaction created successfully",
		"account_id":        transaction.AccountID,
		"transaction_id":    transactionInfo.ID,
		"operation_type_id": transaction.OperationTypeId,
		"amount":            transaction.Amount})
}
