package repository

import "github.com/g-celente/AuthService/src/core/model"

type HistoryEventRepository interface {
	Save(recommendation *model.HistoryEvent) error
}
