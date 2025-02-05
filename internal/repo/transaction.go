package repo

import (
	"github.com/vamshi1997/pismo-assessment/internal/boot"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"log"
)

func CreateTransaction(transaction model.Transaction) (*model.Transaction, error) {
	db := boot.GetDB()
	if err := db.Create(&transaction); err.Error != nil {
		log.Println("Error while creating transaction: ", err.Error)
		return nil, err.Error
	}

	return &transaction, nil
}
