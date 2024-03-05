package models

import (
	"github.com/Quero-Freela/system/server/utils"
	"time"
)

type BidMessage struct {
	utils.BaseModel
	BidID   uint       `json:"bid_id" gorm:"not null;"`
	Bid     *Bid       `json:"-" gorm:"hasOne:bids;not null;foreignKey:BidID"`
	Message string     `json:"message" gorm:"not null;type:text"`
	FromID  uint       `json:"from_id" gorm:"not null;"`
	From    *User      `json:"from" gorm:"not null;hasOne:users;foreignKey:FromID"`
	ToID    uint       `json:"to_id" gorm:"not null;"`
	To      *User      `json:"to" gorm:"hasOne:users;foreignKey:ToID"`
	When    *time.Time `json:"when"`
}
