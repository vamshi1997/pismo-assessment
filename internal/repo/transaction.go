package repo

import (
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"log"
)

func (r *Repository) CreateTransaction(transaction model.Transaction) (*model.Transaction, error) {
	if err := r.db.Create(&transaction); err.Error != nil {
		log.Println("Error while creating transaction: ", err.Error)
		return nil, err.Error
	}

	return &transaction, nil
}
