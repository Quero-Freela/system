package models

import "github.com/Quero-Freela/system/server/utils"

type ProjectImage struct {
	utils.BaseModel
	ProjectID uint    `json:"project_id" gorm:"not null;"`
	Project   Project `json:"-" goem:"hasOne:projects;foreignKey:ProjectID"`
}
