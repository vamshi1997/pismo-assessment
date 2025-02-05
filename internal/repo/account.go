package repo

import (
	"github.com/vamshi1997/pismo-assessment/internal/boot"
	"github.com/vamshi1997/pismo-assessment/internal/model"
	"log"
)

func CreateAccount(account model.Account) (model.Account, error) {
	db := boot.GetDB()
	if err := db.Create(&account); err.Error != nil {
		log.Println("Error while creating account: ", err.Error)
		return account, err.Error
	}

	log.Println("account created successfully")
	return account, nil
}

func GetAccount(accountId uint) (*model.Account, error) {
	db := boot.GetDB()

	var accountInfo model.Account

	if err := db.Where("id = ?", accountId).First(&accountInfo); err.Error != nil {
		log.Println("Error while fetching account info: ", err.Error)
		return nil, err.Error
	}

	return &accountInfo, nil
}
