package routes

import (
	"net/http"

	"github.com/AdityaAWP/IdiomaMate/cmd/api/initialization"
	"github.com/AdityaAWP/IdiomaMate/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, deps *initialization.Dependencies) {
	// Health check
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Public auth routes
	auth := router.Group("/api/v1/auth")
	{
		auth.POST("/register", deps.AuthHandler.Register)
		auth.POST("/login", deps.AuthHandler.Login)
		auth.POST("/refresh", deps.AuthHandler.RefreshToken)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(deps.JWTManager))
	{
		protected.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			email, _ := c.Get("email")
			username, _ := c.Get("username")
			c.JSON(http.StatusOK, gin.H{
				"user_id":  userID,
				"email":    email,
				"username": username,
			})
		})
	}
}
