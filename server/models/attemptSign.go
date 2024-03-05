package models

import "github.com/Quero-Freela/system/server/utils"

type AttemptSign struct {
	utils.BaseModel
	Provider     string `json:"provider" validate:"required" gorm:"not null;size:1024;uniqueIndex:idx_provider_state_request"`
	StateRequest string `json:"state_request" validate:"required" gorm:"not null;size:1024;uniqueIndex:idx_provider_state_request"`
	CallbackURL  string `json:"callback_url" validate:"required" gorm:"not null;size:1024"`
	TimeToLive   int    `json:"time_to_live" validate:"required" gorm:"not null;default:300"`
	Code         string `json:"code" gorm:"size:1024"`
	AccessToken  string `json:"access_token" gorm:"type:text"`
	UserID       *uint  `json:"-"`
	User         *User  `json:"-" gorm:"hasOne:users"`
}
