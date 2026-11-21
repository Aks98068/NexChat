package main

import (
	"log"
	"strings"

	"github.com/Aks98068/NextChat/internal/config"
	"github.com/Aks98068/NextChat/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {

	// loading the configuration form the environment variables
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	if strings.EqualFold(cfg.AppEnv, "production") {
		gin.SetMode(gin.ReleaseMode)
	}

	// Datatabse connection and migration
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("database connection err : %v", err)
	}

	log.Println("Postgresql databse connected successfully")

	if sqlDB, err := db.DB(); err == nil {
		defer sqlDB.Close()
	}

	if err := database.Migratedb(db); err != nil {
		log.Fatalf("database migration failed: %v", err)
	}
	log.Println("Database migration completed successfully")
}
