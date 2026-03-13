package main

import (
	"fmt"
	"net/http"

	"github.com/AdityaAWP/IdiomaMate/cmd/api/migration"
	"github.com/AdityaAWP/IdiomaMate/pkg/config"
	"github.com/AdityaAWP/IdiomaMate/pkg/database"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig(".")
	database.InitDB(cfg)
	migration.Migrate()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	router.Run(addr)
}
