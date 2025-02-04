package model

import (
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	ID             uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	DocumentNumber string `json:"document_number" gorm:"uniqueIndex;not null;type:varchar(255)"`
}
