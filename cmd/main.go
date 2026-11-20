package main

import (
	"log/slog"
	"os"
	"pizza-tracker/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := loadConfig()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	dbModel, err := models.InitDB(cfg.DBPath)
	if err != nil {
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	slog.Info("Database initialized successfully")
	RegisterCustomValidators()

	h := NewHandler(dbModel)
	router := gin.Default()
	if err := loadTemplates(router); err != nil {
		slog.Error("Failed to Load Templates", "error", err)
		os.Exit(1)
	}
	sessionStore := setupSessionStore(dbModel.Order.DB, []byte(cfg.SessionSecretKey))
	setupRoutes(router, h, &sessionStore)
	slog.Info("Starting server", "url", "http://localhost:"+ cfg.Port)
	router.Run(":" + cfg.Port)
}