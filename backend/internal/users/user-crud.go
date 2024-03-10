package users

import (
	"fintrackpro/backend/internal/database"
	"fintrackpro/backend/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

// Register creates a new user in the database with encrypted password.
func Register(userName, email, password string) error {
	user := models.User{
		UserName: userName,
		Email:    email,
	}

	// Encrypt the password.
	err := user.EncryptPassword(password)
	if err != nil {
		zap.L().Sugar().Errorf("Failed to encrypt password: %v", err)
		return err
	}

	// Save the user to the database.
	if err := database.Create(&user); err != nil {
		zap.L().Sugar().Errorf("Failed to create user: %v", err)
		return err
	}

	return nil
}

// Login validates user's credentials.
func Login(email, password string) (*models.User, error) {
	user, err := database.ValidateUser(email, password)
	if err != nil {
		zap.L().Sugar().Errorf("Login failed for user %s: %v", email, err)
		return nil, err
	}

	return user, nil
}

// UpdateUserProfile updates the details of an existing user based on provided input.
func UpdateUserProfile(userID uuid.UUID, input models.UpdateUserInput) error {
	// Fetch the existing user details.
	user, err := database.GetUserByID(userID)
	if err != nil {
		zap.L().Sugar().Errorf("Failed to get user by ID: %v", err)
		return err
	}

	// Update the user details if input fields are non-nil.
	if input.UserName != nil {
		user.UserName = *input.UserName
	}
	if input.Email != nil {
		user.Email = *input.Email
	}

	// Save the updated user details to the database.
	if _, err := database.UpdateUser(userID, *user); err != nil {
		zap.L().Sugar().Errorf("Failed to update user profile: %v", err)
		return err
	}

	return nil
}

// CreateTransaction adds a new transaction for a user.
func CreateTransaction(userID uuid.UUID, title string, amount float64, transactionType, category string, timeStamp time.Time) error {
	transaction := models.Transaction{
		Title:     title,
		Amount:    amount,
		Type:      transactionType,
		Category:  category,
		UserID:    userID,
		TimeStamp: timeStamp,
	}

	if err := database.CreateTransaction(&transaction); err != nil {
		zap.L().Sugar().Errorf("Failed to create transaction: %v", err)
		return err
	}

	return nil
}

// CreateBudget adds a new budget for a user.
func CreateBudget(userID uuid.UUID, title, description string, amount float64) error {
	budget := models.BudgetPlan{
		UserID:            userID,
		BudgetTitle:       title,
		BudgetDescription: description,
		BudgetAmount:      amount,
	}

	if err := database.CreateBudget(&budget); err != nil {
		zap.L().Sugar().Errorf("Failed to create budget: %v", err)
		return err
	}

	return nil
}
