package repo

import (
	"errors"
	"fmt"
	"github.com/vamshi1997/pismo-assessment/internal/boot"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"gorm.io/gorm"
	"log"
)

func CreateAccount(account model.Account) {
	db := boot.GetDB()

	if err := db.Where("document_number = ?", account.DocumentNumber).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			db.Create(&account)
			fmt.Println("Account created successfully")
		} else {
			fmt.Println("Error while creating account in database:", err)
		}
	} else {
		// Video already exists
		fmt.Println("Account already exists for Document Number:", account.DocumentNumber)
	}
}

func GetAccount(accountId int) (*model.Account, error) {
	db := boot.GetDB()

	var accountInfo model.Account

	if err := db.Where("id = ?", accountId).First(&accountInfo); err.Error != nil {
		log.Println("Error while fetching account info: ", err.Error)
		return nil, nil
	}

	return &accountInfo, nil
}

func CreateTransaction(transaction model.Transaction) error {
	db := boot.GetDB()
	if err := db.Create(&transaction); err.Error != nil {
		log.Println("Error while creating transaction: ", err.Error)
		return err.Error
	}
	return nil
}
