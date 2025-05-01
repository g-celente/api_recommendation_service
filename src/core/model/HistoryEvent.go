package model

import (
	"github.com/lib/pq"
)

type HistoryEvent struct {
	ID      uint   `gorm:"primaryKey"`
	UserID  string `gorm:"index"`
	EventID string
	Tags    pq.StringArray `gorm:"type:text[]" json:"tags"` // PostgreSQL array
}
