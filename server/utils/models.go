package utils

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type NullUInt sql.NullInt64

// Scan implements the Scanner interface.
func (n *NullUInt) Scan(value any) error {
	ni64 := &sql.NullInt64{
		Int64: n.Int64,
		Valid: n.Valid,
	}

	// check if value is uint
	if v, ok := value.(uint); ok {
		value = int64(v)
	}

	err := ni64.Scan(value)

	if err == nil {
		n.Int64 = ni64.Int64
		n.Valid = ni64.Valid
	}

	return err
}

type BaseModel struct {
	ID        uint           `json:"id" gorm:"primarykey;autoIncrement;not null;type:bigint"`
	CreatedAt time.Time      `json:"created_at" gorm:"not null;index;<-:create"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"not null;index;<-:create;<-:update;default:NOW()"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"null;index;<-:update"`
	CreatedBy uint           `json:"created_by" gorm:"not null;index;<-:create;type:bigint"`
	UpdatedBy uint           `json:"updated_by" gorm:"not null;index;<-:create;<-:update;type:bigint"`
	DeletedBy *NullUInt      `json:"deleted_by" gorm:"null;index;<-:update;type:bigint"`
	Deleted   bool           `json:"deleted" gorm:"not null;->;type:boolean GENERATED ALWAYS AS (deleted_at IS NOT NULL) STORED"`
}
