package models

import (
	"database/sql"
	"github.com/Quero-Freela/system/server/exceptions"
	"github.com/Quero-Freela/system/server/utils"
	"gorm.io/gorm"
)

type User struct {
	utils.BaseModel
	Name       string         `json:"name" validate:"required" gorm:"not null;size:1024"`
	Email      string         `json:"email" validate:"required,email" gorm:"not null;unique;size:1024;"`
	VerifiedAt sql.NullTime   `json:"verified_at" gorm:"type:timestamp"`
	Verified   bool           `json:"verified" gorm:"not null;->;type:boolean GENERATED ALWAYS AS (verified_at IS NOT NULL) STORED"`
	Roles      []*Role        `json:"roles" gorm:"many2many:user_roles;"`
	Attempts   []*AttemptSign `json:"attempts" gorm:"hasMany:attempt_sings"`
	JwtSecret  []byte         `json:"-" gorm:"not null"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hasEmail := tx.Exec("SELECT 1 AS has_email FROM users WHERE deleted = false AND email = ?", u.Email)
	if hasEmail.Error != nil {
		return hasEmail.Error
	}

	if hasEmail.RowsAffected > 0 {
		return exceptions.NewConflictError("there is already a user with this email")
	}

	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	hasEmail := tx.Exec("SELECT 1 AS has_email FROM users WHERE deleted = false AND email = ? AND id != ?", u.Email, u.ID)
	if hasEmail.Error != nil {
		return hasEmail.Error
	}

	if hasEmail.RowsAffected > 0 {
		return exceptions.NewConflictError("there is already a user with this email")
	}

	return nil
}
