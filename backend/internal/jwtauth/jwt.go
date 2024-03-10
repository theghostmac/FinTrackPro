package jwtauth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims struct to contain the JWT claims.
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.StandardClaims
}

// GenerateUserToken generates a JWT token for the user.
func GenerateUserToken(userID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := &Claims{
		UserID: userID,
		Role:   "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

// ValidateToken validates a given JWT token and returns the user UserID and role.
func ValidateToken(tokenString string) (*uuid.UUID, string, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, "", err
	}

	if !token.Valid {
		return nil, "", err
	}

	return &claims.UserID, claims.Role, nil
}
