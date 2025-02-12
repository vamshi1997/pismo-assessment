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

	tests := []struct {
		name           string
		input          model.Transaction
		mockBehavior   func(*mock.MockIRepository)
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "Valid Purchase Transaction",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 1,
				Amount:          -100.0,
			},
			mockBehavior: func(m *mock.MockIRepository) {
				m.EXPECT().
					GetAccount(uint(1)).
					Return(&model.Account{ID: 1, DocumentNumber: "12345678901"}, nil)

				m.EXPECT().
					CreateTransaction(gomock.Any()).
					Return(&model.Transaction{
						ID:              1,
						AccountID:       1,
						OperationTypeId: 1,
						Amount:          -100.0,
						Balance:         -100.0,
					}, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg":               "transaction created successfully",
				"account_id":        float64(1),
				"transaction_id":    float64(1),
				"operation_type_id": float64(1),
				"amount":            float64(-100.0),
			},
		},
		{
			name: "Invalid Operation Type",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 10, // Invalid operation type
				Amount:          -100.0,
			},
			mockBehavior:   func(m *mock.MockIRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error_msg": "Invalid operation type",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Invalid Amount for Purchase (Positive Amount)",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 1,
				Amount:          100.0, // Should be negative for purchase
			},
			mockBehavior:   func(m *mock.MockIRepository) {},
			expectedStatus: http.StatusBadRequest,
			expectedBody: map[string]interface{}{
				"error_msg": "Amount can not be positive for this operation type",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Invalid Account",
			input: model.Transaction{
				AccountID:       999,
				OperationTypeId: 1,
				Amount:          -100.0,
			},
			mockBehavior: func(m *mock.MockIRepository) {
				m.EXPECT().
					GetAccount(uint(999)).
					Return((*model.Account)(nil), errors.New("account not found"))
			},
			expectedStatus: http.StatusNotFound,
			expectedBody: map[string]interface{}{
				"error":     "account not found",
				"error_msg": "Account not found",
				"msg":       "Not able to create transaction",
			},
		},
		{
			name: "Valid Credit Voucher Transaction",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 4,
				Amount:          100.0,
			},
			mockBehavior: func(m *mock.MockIRepository) {
				// Create expected calls in sequence
				gomock.InOrder(
					// 1. Check account exists
					m.EXPECT().
						GetAccount(uint(1)).
						Return(&model.Account{ID: 1, DocumentNumber: "12345678901"}, nil),

					// 2. Get previous transactions
					m.EXPECT().
						GetPreviousTransactions().
						Return([]model.Transaction{
							{
								ID:              1,
								AccountID:       1,
								OperationTypeId: 1,
								Amount:          -150.0,
								Balance:         -150.0,
							},
						}, nil),

					// 3. Update transaction balance
					m.EXPECT().
						UpdateTransactionBalance(gomock.Any(), uint(1)).
						Return(&model.Transaction{
							ID:              1,
							AccountID:       1,
							OperationTypeId: 1,
							Amount:          -150.0,
							Balance:         0.0,
						}, nil),

					// 4. Create new transaction
					m.EXPECT().
						CreateTransaction(gomock.Any()).
						Return(&model.Transaction{
							ID:              2,
							AccountID:       1,
							OperationTypeId: 4,
							Amount:          100.0,
							Balance:         0.0,
						}, nil),
				)
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg":               "transaction created successfully",
				"account_id":        float64(1),
				"transaction_id":    float64(2),
				"operation_type_id": float64(4),
				"amount":            float64(100.0),
			},
		},
		{
			name: "Credit Voucher With No Previous Transactions",
			input: model.Transaction{
				AccountID:       1,
				OperationTypeId: 4,
				Amount:          100.0,
			},
			mockBehavior: func(m *mock.MockIRepository) {
				// First, expect GetAccount call
				m.EXPECT().
					GetAccount(uint(1)).
					Return(&model.Account{ID: 1, DocumentNumber: "12345678901"}, nil)

				// Then, expect GetPreviousTransactions call returning empty slice
				m.EXPECT().
					GetPreviousTransactions().
					Return([]model.Transaction{}, nil)

				// Expect CreateTransaction call
				m.EXPECT().
					CreateTransaction(gomock.Any()).
					DoAndReturn(func(transaction model.Transaction) (*model.Transaction, error) {
						return &model.Transaction{
							ID:              1,
							AccountID:       1,
							OperationTypeId: 4,
							Amount:          100.0,
							Balance:         100.0,
						}, nil
					})
			},
			expectedStatus: http.StatusOK,
			expectedBody: map[string]interface{}{
				"msg":               "transaction created successfully",
				"account_id":        float64(1),
				"transaction_id":    float64(1),
				"operation_type_id": float64(4),
				"amount":            float64(100.0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock.NewMockIRepository(ctrl)
			tt.mockBehavior(mockRepo)
			controller := NewController(mockRepo)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Create request body
			jsonBytes, _ := json.Marshal(tt.input)
			c.Request = httptest.NewRequest(http.MethodPost, "/transactions", bytes.NewBuffer(jsonBytes))
			c.Request.Header.Set("Content-Type", "application/json")

			// Execute
			controller.CreateTransaction(c)

			// Assert
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
