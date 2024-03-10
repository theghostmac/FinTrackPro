package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID       uuid.UUID     `gorm:"type:uuid;primaryKey"`
	UserName     string        `gorm:"not null;unique"`
	Email        string        `gorm:"not null;unique"`
	PasswordHash string        `gorm:"not null"` // Store hashed password
	Transactions []Transaction `gorm:"foreignKey:UserID"`
	CreatedAt    time.Time     `gorm:"not null"`
}

type UpdateUserInput struct {
	UserName *string `json:"userName,omitempty"`
	Email    *string `json:"email,omitempty"`
}

// BeforeCreate is a GORM hook for UUID generation
func (user *User) BeforeCreate(_ *gorm.DB) (err error) {
	user.UserID = uuid.New()
	return
}

// EncryptPassword hashes the password using bcrypt.
func (user *User) EncryptPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PasswordHash = string(hashedPassword)
	return nil
}

// ComparePasswordAndHash checks if the provided password matches the hashed password.
func (user *User) ComparePasswordAndHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}
