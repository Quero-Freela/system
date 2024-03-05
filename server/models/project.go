package models

import (
	"github.com/Quero-Freela/system/server/utils"
	"time"
)

type Project struct {
	utils.BaseModel
	Name         string         `json:"name" gorm:"not null;type:varchar(64)"`
	Description  string         `json:"description" gorm:"not null;type:text"`
	Images       []ProjectImage `json:"images" gorm:"hasMany:project_images;foreignKey:ProjectID"`
	Bids         []*Bid         `json:"bids" gorm:"hasMany:bids;foreignKey:ProjectID"`
	Budget       float64        `json:"budget" gorm:"not null;type:decimal(10,2)"`
	Deadline     int            `json:"deadline" gorm:"not null;default:7"`
	DeadlineType DeadlineType   `json:"deadline_type" gorm:"varchar(6);not null;default:days"`
	PublishedOn  *time.Time     `json:"published_on"`
	Published    bool           `json:"published" gorm:"not null;->;type:boolean GENERATED ALWAYS AS (published_on IS NOT NULL) STORED"`
	Active       bool           `json:"active" gorm:"not null;default:true;"`
	Categories   []*Category    `json:"categories" gorm:"many2many:project_categories;"`
}
