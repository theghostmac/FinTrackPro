package database

import (
	"errors"
	"fintrackpro/backend/internal/models"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Create(user *models.User) error {
	if DB == nil {
		zap.L().Sugar().Error("DB connection is nil")
		return errors.New("DB connection is nil")
	}
	return DB.Create(user).Error
}

// GetUserByID returns a user by its ID.
func GetUserByID(userID uuid.UUID) (*models.User, error) {
	var user models.User
	if err := DB.First(user, userID).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func ValidateUser(email, password string) (*models.User, error) {
	var user models.User
	err := DB.Where("email =?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	// Compare the provided password with the stored hash.
	if !user.ComparePasswordAndHash(password) {
		return nil, gorm.ErrRecordNotFound
	}

	return &user, nil
}

func UpdateUser(userID uuid.UUID, updatedUser models.User) (*models.User, error) {
	var user models.User
	if err := DB.First(&user, "user_id = ?", userID).Error; err != nil {
		return nil, err
	}

	// Update required fields.
	user.UserName = updatedUser.UserName
	user.Email = updatedUser.Email

	if err := DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
