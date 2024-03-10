package users

import (
	"fintrackpro/backend/internal/database"
	"fintrackpro/backend/internal/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// RegisterHandler handles user registration requests.
func RegisterHandler(c *gin.Context) {
	var request struct {
		UserName string `json:"userName"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := Register(request.UserName, request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// LoginHandler handles user login requests.
func LoginHandler(c *gin.Context) {
	var request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := Login(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "user": user})
}

func UpdateUserProfileHandler(c *gin.Context) {
	var input models.UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := getUserIDFromContext(c)

	// Pass input to the update function
	if err := UpdateUserProfile(userID, input); err != nil {
		zap.L().Sugar().Errorf("Failed to update user profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User profile updated successfully"})
}

func GetMyProfileHandler(c *gin.Context) {
	userID := getUserIDFromContext(c)

	user, err := database.GetUserByID(userID)
	if err != nil {
		zap.L().Sugar().Errorf("Failed to get user profile: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user profile"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func CreateTransactionHandler(c *gin.Context) {
	var request struct {
		Title     string    `json:"title"`
		Amount    float64   `json:"amount"`
		Type      string    `json:"type"`
		Category  string    `json:"category"`
		TimeStamp time.Time `json:"timeStamp"`
	}
	userID := getUserIDFromContext(c)

	if err := CreateTransaction(userID, request.Title, request.Amount, request.Type, request.Category, request.TimeStamp); err != nil {
		zap.L().Sugar().Errorf("Failed to create transaction: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction created successfully"})
}

func CreateBudgetHandler(c *gin.Context) {
	var request struct {
		Title       string  `json:"title"`
		Description string  `json:"description"`
		Amount      float64 `json:"amount"`
	}
	userID := getUserIDFromContext(c)

	if err := CreateBudget(userID, request.Title, request.Description, request.Amount); err != nil {
		zap.L().Sugar().Errorf("Failed to create budget: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create budget"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Budget created successfully"})
}
