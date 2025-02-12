package repo

import (
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

type IRepository interface {
	CreateAccount(account model.Account) (model.Account, error)
	GetAccount(accountId uint) (*model.Account, error)
	CreateTransaction(transaction model.Transaction) (*model.Transaction, error)
	GetPreviousTransactions() ([]model.Transaction, error)
	UpdateTransactionBalance(balance float64, transactionId uint) (*model.Transaction, error)
}

func NewRepository(db *gorm.DB) IRepository {
	return &Repository{
		db: db,
	}
}
