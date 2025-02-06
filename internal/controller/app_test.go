package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"github.com/vamshi1997/pismo-assessment/internal/repo/mock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	Status(c)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.NoError(t, err)
	assert.Equal(t, "ok", response["status"])
}

func TestController_CreateAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	controller := NewController(mockRepo)

	tests := []struct {
		name           string
		input          model.Account
		mockBehavior   func(mock *mock.MockIRepository, account model.Account)
		expectedStatus int
		expectedBody   gin.H
	}{
		{
			name: "Success",
			input: model.Account{
				DocumentNumber: "12345678901",
			},
			mockBehavior: func(mock *mock.MockIRepository, account model.Account) {
				mock.EXPECT().
					CreateAccount(account).
					Return(model.Account{ID: 1, DocumentNumber: account.DocumentNumber}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: gin.H{
				"document_number": "12345678901",
				"account_id":      float64(1),
				"msg":             "Account created successfully",
			},
		},
		{
			name: "Invalid Document Number",
			input: model.Account{
				DocumentNumber: "123",
			},
			mockBehavior:   func(mock *mock.MockIRepository, account model.Account) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: gin.H{
				"error_msg": "Document number given is not valid",
				"msg":       "Not able to create account",
			},
		},
		{
			name: "Database Error",
			input: model.Account{
				DocumentNumber: "12345678901",
			},
			mockBehavior: func(mock *mock.MockIRepository, account model.Account) {
				mock.EXPECT().
					CreateAccount(account).
					Return(model.Account{}, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: gin.H{
				"error":     "database error",
				"error_msg": "Internal Server Error",
				"msg":       "Not able to create account",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonValue, _ := json.Marshal(tt.input)
			c.Request = httptest.NewRequest("POST", "/accounts", bytes.NewBuffer(jsonValue))
			c.Request.Header.Set("Content-Type", "application/json")

			tt.mockBehavior(mockRepo, tt.input)

			controller.CreateAccount(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			for key, expectedValue := range tt.expectedBody {
				assert.Equal(t, expectedValue, response[key])
			}
		})
	}
}

func TestController_GetAccount(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	controller := NewController(mockRepo)

	tests := []struct {
		name           string
		accountID      string
		mockBehavior   func(mock *mock.MockIRepository, accountID uint)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name:      "Success",
			accountID: "1",
			mockBehavior: func(mock *mock.MockIRepository, accountID uint) {
				mock.EXPECT().
					GetAccount(accountID).
					Return(&model.Account{
						ID:             1,
						DocumentNumber: "12345678901",
					}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"account_id":      float64(1),
				"document_number": "12345678901",
				"msg":             "Account details fetched successfully",
			},
		},
		{
			name:           "Invalid Account ID Format",
			accountID:      "invalid",
			mockBehavior:   func(mock *mock.MockIRepository, accountID uint) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error_msg": "Not valid accountId",
				"msg":       "Not able to fetch account details",
			},
		},
		{
			name:      "Account Not Found",
			accountID: "3",
			mockBehavior: func(mock *mock.MockIRepository, accountID uint) {
				mock.EXPECT().
					GetAccount(accountID).
					Return(nil, errors.New("account not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error":     "account not found",
				"error_msg": "Account not found",
				"msg":       "Not able to fetch account details",
			},
		},
		{
			name:      "Empty Account Details",
			accountID: "4",
			mockBehavior: func(mock *mock.MockIRepository, accountID uint) {
				mock.EXPECT().
					GetAccount(accountID).
					Return(&model.Account{}, nil)
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error_msg": "Account Details are empty",
				"msg":       "Not able to fetch account details",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			c.Request = httptest.NewRequest("GET", "/accounts/"+tt.accountID, nil)
			c.Params = []gin.Param{{Key: "accountId", Value: tt.accountID}}

			// Set mock behavior
			if accountID, err := strconv.Atoi(tt.accountID); err == nil {
				tt.mockBehavior(mockRepo, uint(accountID))
			}

			controller.GetAccount(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			for key, expectedValue := range tt.expectedBody {
				assert.Equal(t, expectedValue, response[key], "mismatch in field: %s", key)
			}
		})
	}
}

func TestController_CreateTransaction(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIRepository(ctrl)
	controller := NewController(mockRepo)

	tests := []struct {
		name           string
		input          model.Transaction
		mockBehavior   func(mock *mock.MockIRepository)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Success - Purchase Transaction",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 1,
				Amount:          -100.00,
			},
			mockBehavior: func(mock *mock.MockIRepository) {
				mock.EXPECT().
					GetAccount(uint(1)).
					Return(&model.Account{ID: 1, DocumentNumber: "12345678901"}, nil)

				mock.EXPECT().
					CreateTransaction(gomock.Any()).
					Return(&model.Transaction{
						ID:              1,
						AccountID:       1,
						OperationTypeId: 1,
						Amount:          -100.00,
					}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg":               "transaction created successfully",
				"account_id":        float64(1),
				"transaction_id":    float64(1),
				"operation_type_id": float64(1),
				"amount":            float64(-100.00),
			},
		},
		{
			name: "Invalid Operation Type",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 999, // Invalid operation type
				Amount:          -100.00,
			},
			mockBehavior:   func(mock *mock.MockIRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error_msg": "Invalid operation type",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Invalid Amount for Purchase Operation",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 1,
				Amount:          100.00, // Positive amount for purchase
			},
			mockBehavior:   func(mock *mock.MockIRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error_msg": "Amount can not be positive for this operation type",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Invalid Amount for Credit Operation",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 4,
				Amount:          -100.00,
			},
			mockBehavior:   func(mock *mock.MockIRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error_mgs": "Amount can not be negative for this operation type",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Account Not Found",
			input: model.Transaction{
				AccountID:       999,
				OperationTypeId: 1,
				Amount:          -100.00,
			},
			mockBehavior: func(mock *mock.MockIRepository) {
				mock.EXPECT().
					GetAccount(uint(999)).
					Return(nil, errors.New("account not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error":     "account not found",
				"error_msg": "Account not found",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Transaction Creation Error",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 1,
				Amount:          -100.00,
			},
			mockBehavior: func(mock *mock.MockIRepository) {
				mock.EXPECT().
					GetAccount(uint(1)).
					Return(&model.Account{ID: 1, DocumentNumber: "12345678901"}, nil)

				mock.EXPECT().
					CreateTransaction(gomock.Any()).
					Return(nil, errors.New("database error"))
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody: map[string]interface{}{
				"error":     "database error",
				"error_msg": "Invalid transaction",
				"msg":       "Not able to create transaction",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			jsonValue, err := json.Marshal(tt.input)
			assert.NoError(t, err)

			c.Request = httptest.NewRequest("POST", "/transactions", bytes.NewBuffer(jsonValue))
			c.Request.Header.Set("Content-Type", "application/json")

			tt.mockBehavior(mockRepo)

			controller.CreateTransaction(c)

			assert.Equal(t, tt.expectedStatus, w.Code)

			var response map[string]interface{}
			err = json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)

			for key, expected := range tt.expectedBody {
				actual, exists := response[key]
				assert.True(t, exists, "Expected key %s not found in response", key)
				assert.Equal(t, expected, actual, "Value mismatch for key %s", key)
			}
		})
	}
}
