package initialization

import (
	"github.com/AdityaAWP/IdiomaMate/internal/handler"
	"github.com/AdityaAWP/IdiomaMate/internal/repository"
	"github.com/AdityaAWP/IdiomaMate/internal/service"
	"github.com/AdityaAWP/IdiomaMate/pkg/config"
	"github.com/AdityaAWP/IdiomaMate/pkg/utils"
	"gorm.io/gorm"
)

type Dependencies struct {
	JWTManager  *utils.JWTManager
	AuthHandler *handler.AuthHandler
}

func InitDependencies(cfg *config.Config, db *gorm.DB) *Dependencies {
	// JWT
	jwtManager := utils.NewJWTManager(
		cfg.JWT.Secret,
		cfg.JWT.AccessTokenExpiration,
		cfg.JWT.RefreshTokenExpiration,
	)

	// Repository
	userRepo := repository.NewUserRepository(db)

	// Service
	authService := service.NewAuthService(userRepo, jwtManager)

	// Handler
	authHandler := handler.NewAuthHandler(authService)

	return &Dependencies{
		JWTManager:  jwtManager,
		AuthHandler: authHandler,
	}
}
