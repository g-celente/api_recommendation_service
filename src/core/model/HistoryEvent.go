package model

import (
	"time"
)

type HistoryEvent struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	EmailSubscriber string    `gorm:"index" json:"emailSubscriber"`
	EventID         uint      `json:"eventId"`
	EventName       string    `json:"eventName"`
	EventTheme      string    `json:"eventTheme"`
	SubscribedAt    time.Time `json:"subscribedAt"` // opcional: registrar quando o evento foi armazenado
}
