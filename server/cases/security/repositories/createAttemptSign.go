package repositories

import (
	"github.com/Quero-Freela/system/server/database"
	"github.com/Quero-Freela/system/server/models"
	"github.com/Quero-Freela/system/server/utils"
)

func CreateAttemptSign(provider, callbackURL string) (*models.AttemptSign, error) {
	state, err := utils.GenerateRandomString(32)

	if err != nil {
		return nil, err
	}

	attempt := &models.AttemptSign{
		Provider:     provider,
		StateRequest: state,
		CallbackURL:  callbackURL,
		TimeToLive:   300,
	}

	db := database.Connect()
	err = database.MigrateAll(db)

	if err != nil {
		return nil, err
	}

	err = db.Create(attempt).Error

	if err != nil {
		return nil, err
	}

	return attempt, nil
}
