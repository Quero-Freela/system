package repositories

import (
	"database/sql"
	"errors"
	"github.com/Quero-Freela/system/server/database"
	"github.com/Quero-Freela/system/server/exceptions"
	"github.com/Quero-Freela/system/server/models"
	"time"
)

func FindAttemptByState(provider, state string) (*models.AttemptSign, error) {
	db := database.Connect()
	err := database.MigrateAll(db)
	if err != nil {
		return nil, err
	}

	attempt := models.AttemptSign{
		Provider:     provider,
		StateRequest: state,
	}

	if err = db.Model(&models.AttemptSign{}).First(&attempt).Error; err == nil && attempt.ID != 0 {
		now := time.Now()
		ttl := attempt.CreatedAt.Add(time.Duration(attempt.TimeToLive) * time.Second)
		if now.After(ttl) && attempt.Code == "" {
			return nil, exceptions.NewNotFoundError("login attempt expired")
		}

		return &attempt, nil
	}

	if err == nil || errors.Is(err, sql.ErrNoRows) {
		return nil, exceptions.NewNotFoundError("invalid login attempt")
	}

	return nil, err
}
