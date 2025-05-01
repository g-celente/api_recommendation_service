package database

import (
	"github.com/g-celente/AuthService/src/core/model"
	"gorm.io/gorm"
)

type HistoryEventPG struct {
	DB *gorm.DB
}

func NewHistoryEventPG(db *gorm.DB) *HistoryEventPG {
	return &HistoryEventPG{DB: db}
}

func (r *HistoryEventPG) Save(rec *model.HistoryEvent) error {
	return r.DB.Create(rec).Error
}
