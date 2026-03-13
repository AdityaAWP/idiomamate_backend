package main

import (
	"fmt"

	"github.com/AdityaAWP/IdiomaMate/cmd/api/initialization"
	"github.com/AdityaAWP/IdiomaMate/cmd/api/migration"
	"github.com/AdityaAWP/IdiomaMate/cmd/api/routes"
	"github.com/AdityaAWP/IdiomaMate/pkg/config"
	"github.com/AdityaAWP/IdiomaMate/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(".")
	database.InitDB(cfg)
	migration.Migrate()

	deps := initialization.InitDependencies(cfg, database.DB)

	router := gin.Default()
	routes.SetupRoutes(router, deps)

	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	router.Run(addr)
}
