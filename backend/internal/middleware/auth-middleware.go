package middleware

import (
	"fintrackpro/backend/internal/jwtauth"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

var logger, _ = zap.NewDevelopment()

func UserAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			logger.Error("Missing Authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing Authorization header"})
			c.Abort()
			return
		}

		prefix := "Bearer "
		if !strings.HasPrefix(authHeader, prefix) {
			logger.Error("Invalid Authorization header format")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, prefix)
		userID, role, err := jwtauth.ValidateToken(tokenString)
		if err != nil {
			logger.Error("Token validation failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		logger.Info("Token validated successfully", zap.Any("userID", *userID), zap.String("role", role))

		// Check if the user is a regular user
		if role != "user" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Access restricted to users"})
			c.Abort()
			return
		}

		// Set both userID and role in the context
		c.Set("userID", *userID)
		c.Set("role", role)

		c.Next()
	}
}
