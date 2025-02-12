package repo

import (
	"fmt"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"log"
	"time"
)

var IST = time.FixedZone("IST", 5*3600+1800)

func (r *Repository) CreateTransaction(transaction model.Transaction) (*model.Transaction, error) {
	if err := r.db.Create(&transaction); err.Error != nil {
		log.Println("Error while creating transaction: ", err.Error)
		return nil, err.Error
	}

	return &transaction, nil
}

func (r *Repository) UpdateTransactionBalance(balance float64, transactionId uint) (*model.Transaction, error) {
	result := r.db.Model(&model.Transaction{}).
		Where("id = ?", transactionId).
		Update("balance", balance)

	if result.Error != nil {
		log.Printf("Error updating transaction balance: %v", result.Error)
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no transaction found with ID: %d", transactionId)
	}

	// Fetch the updated transaction
	var updatedTransaction model.Transaction
	if err := r.db.First(&updatedTransaction, transactionId).Error; err != nil {
		return nil, fmt.Errorf("error fetching updated transaction: %w", err)
	}

	return &updatedTransaction, nil
}

func (r *Repository) GetPreviousTransactions() ([]model.Transaction, error) {
	var transactions []model.Transaction
	currentDate := time.Now().In(IST)

	result := r.db.
		Where("created_at < ?", currentDate).
		Where("balance < ?", 0).
		Order("event_date ASC").
		Find(&transactions)

	if result.Error != nil {
		log.Printf("Error while fetching previous transactions: %v", result.Error)
		return nil, result.Error
	}

	return transactions, nil
}
