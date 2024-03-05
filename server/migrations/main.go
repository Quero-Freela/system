//go:build !lambda && amd64

// nolint
package main

import (
	"context"
	"github.com/Quero-Freela/system/server/database"
	"github.com/Quero-Freela/system/server/models"
	"github.com/joho/godotenv"
	"google.golang.org/appengine/log"
)

func init() {
	envPath := ".env"
	maxRetry := 3
	var err error
	for err = godotenv.Load(envPath); err != nil && maxRetry > 0; err = godotenv.Load(envPath) {
		maxRetry--
		envPath = "../" + envPath
	}

	if err != nil {
		panic(err)
	}
}

func main() {
	db := database.Connect()

	mods := []interface{}{
		&models.User{},
		&models.AttemptSign{},
		&models.Role{},
		&models.Project{},
		&models.ProjectImage{},
		&models.Bid{},
		&models.BidMessage{},
	}

	err := db.Migrator().AutoMigrate(mods...)

	if err != nil {
		log.Errorf(context.Background(), "Error migrating models: %v", err)
	}
}
