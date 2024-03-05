package models

import (
	"github.com/Quero-Freela/system/server/utils"
	"time"
)

type Bid struct {
	utils.BaseModel
	ProjectID    uint          `json:"project_id" gorm:"not null;"`
	Project      *Project      `json:"-" gorm:"hasOne:projects;not null;foreignKey:ProjectID"`
	BidderID     uint          `json:"bidder_id" gorm:"not null;"`
	Bidder       *User         `json:"bidder" gorm:"not null;hasOne:users;foreignKey:BidderID"`
	Price        float64       `json:"price" gorm:"not null;type:decimal(10,2)"`
	Deadline     int           `json:"deadline" gorm:"not null;default:7"`
	DeadlineType DeadlineType  `json:"deadline_type" gorm:"varchar(6);not null;default:days"`
	Accepted     bool          `json:"accepted" gorm:"not null;default:false"`
	Messages     []*BidMessage `json:"messages" gorm:"hasMany:bid_messages;foreignKey:BidID"`
	StartOn      *time.Time    `json:"start_on"`
}
