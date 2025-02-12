package model

import (
	"gorm.io/gorm"
	"time"
)

var IST = time.FixedZone("IST", 5*3600+1800)

type Transaction struct {
	gorm.Model
	ID              uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	AccountID       uint    `json:"account_id" gorm:"not null;"`
	Amount          float64 `json:"amount" gorm:"not null;"`
	Balance         float64 `json:"balance" gorm:"not null;"`
	OperationTypeId uint    `json:"operation_type_id" gorm:"not null;"`
	EventDate       string  `json:"event_date" gorm:"not null;type:timestamp(6);"`
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.EventDate = time.Now().In(IST).Format("2006-01-02T15:04:05.999999")
	return
}
