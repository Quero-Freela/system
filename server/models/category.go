package models

import "github.com/Quero-Freela/system/server/utils"

type Category struct {
	utils.BaseModel
	Name           string      `json:"name" gorm:"not null;size:1024"`
	Description    string      `json:"description" gorm:"not null;type:text"`
	ParentID       *uint       `json:"parent_id" gorm:"null;"`
	ParentCategory *Category   `json:"parent_category" gorm:"hasOne:categories;foreignKey:ParentID;null;"`
	SubCategories  []*Category `json:"sub_categories" gorm:"hasMany:categories;foreignKey:ParentID;"`
}
