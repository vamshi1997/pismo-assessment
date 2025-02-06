package repo

import (
	"log"

	"github.com/vamshi1997/pismo-assessment/internal/model"
)

func (r *Repository) CreateAccount(account model.Account) (model.Account, error) {
	if err := r.db.Create(&account); err.Error != nil {
		log.Println("Error while creating account: ", err.Error)
		return account, err.Error
	}

	log.Println("account created successfully")
	return account, nil
}

func (r *Repository) GetAccount(accountId uint) (*model.Account, error) {
	var accountInfo model.Account

	if err := r.db.Where("id = ?", accountId).First(&accountInfo); err.Error != nil {
		log.Println("Error while fetching account info: ", err.Error)
		return nil, err.Error
	}

	return &accountInfo, nil
}
