package model

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	ID              uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID       uint    `json:"account_id" gorm:"uniqueIndex;not null;"`
	Amount          float64 `json:"amount" gorm:"not null;"`
	OperationTypeId uint    `json:"operation_type_id" gorm:"not null;"`
	EventDate       string  `json:"event_date" gorm:"not null;type:timestamp(6);"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.EventDate = time.Now().UTC().Format("2006-01-02T15:04:05.999999") // Use UTC for consistency
	return
}
