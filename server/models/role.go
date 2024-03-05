package models

import (
	"github.com/Quero-Freela/system/server/utils"
)

type Role struct {
	utils.BaseModel
	Name  string  `json:"name" validate:"required" gorm:"not null;size:32"`
	Users []*User `gorm:"many2many:user_roles;"`
}
