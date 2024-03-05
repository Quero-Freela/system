package repositories

import (
	"github.com/Quero-Freela/system/server/database"
	models2 "github.com/Quero-Freela/system/server/models"
)

func UpdateAttemptUserToken(attempt *models2.AttemptSign, user *models2.User, code, token string) error {
	db := database.Connect()
	err := database.MigrateAll(db)
	if err != nil {
		return err
	}

	attempt.User = user
	attempt.Code = code
	attempt.AccessToken = token

	return db.Save(attempt).Error
}
