package api

import (
	"fintrackpro/backend/internal/users"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// CORS middleware configuration
	config := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}

	router.Use(cors.New(config))

	// Server health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "UP"})
	})

	// Custom 404 error page
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// User routes
	router.POST("/register", users.RegisterHandler)
	router.POST("/login", users.LoginHandler)
	router.PUT("/user", users.UpdateUserProfileHandler)
	router.POST("/transaction", users.CreateTransactionHandler)
	router.POST("/budget", users.CreateBudgetHandler)

	return router
}
