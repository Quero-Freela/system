package repositories

import (
	"database/sql"
	"errors"
	"github.com/Quero-Freela/system/server/database"
	"github.com/Quero-Freela/system/server/exceptions"
	"github.com/Quero-Freela/system/server/models"
)

func FindUserByEmail(email string) (*models.User, error) {
	db := database.Connect()
	err := database.MigrateAll(db)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Email: email,
	}

	if err = db.Model(&models.User{}).Preload("Roles").First(&user).Error; err == nil && user.ID != 0 {
		return &user, nil
	}

	if err == nil || errors.Is(err, sql.ErrNoRows) {
		return nil, exceptions.NewNotFoundError("User not found")
	}

	return nil, err
}
