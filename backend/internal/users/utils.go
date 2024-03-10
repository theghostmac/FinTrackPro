package users

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getUserIDFromContext(c *gin.Context) uuid.UUID {
	userID, exists := c.Get("userID")
	if !exists {
		return uuid.Nil // Consider how you want to handle this case.
	}

	// Convert interface{} to uuid.UUID
	uuidUserID, ok := userID.(uuid.UUID)
	if !ok {
		return uuid.Nil // Consider logging this anomaly.
	}

	return uuidUserID
}
