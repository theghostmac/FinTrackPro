package database

import (
	"fintrackpro/backend/internal/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	// BeforeCreate hook in the model will automatically be called to generate UUID
	err := DB.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateBudget(budgetPlan *models.BudgetPlan) error {
	// BeforeCreate hook in the model will automatically be called to generate UUID
	err := DB.Create(budgetPlan).Error
	if err != nil {
		return err
	}
	return nil
}
