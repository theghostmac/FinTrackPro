package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	TxID      uuid.UUID `gorm:"type:uuid;primaryKey"`
	Title     string
	Amount    float64
	Type      string `gorm:"not null"` // "income" or "expense"
	Category  string
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	TimeStamp time.Time `gorm:"not null"`
}

// BeforeCreate hook used to generate a new UUID before inserting a new record.
func (transaction *Transaction) BeforeCreate(_ *gorm.DB) (err error) {
	transaction.TxID = uuid.New()
	return
}

type BudgetPlan struct {
	BudgetID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID            uuid.UUID `gorm:"type:uuid;not null"`
	BudgetTitle       string    `gorm:"not null"`
	BudgetDescription string
	BudgetAmount      float64 `gorm:"not null"`
}

// BeforeCreate hook used to generate a new UUID before inserting a new record.
func (budgetPlan *BudgetPlan) BeforeCreate(_ *gorm.DB) (err error) {
	budgetPlan.BudgetID = uuid.New()
	return
}
