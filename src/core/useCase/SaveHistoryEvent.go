package useCase

import (
	"github.com/g-celente/AuthService/src/core/model"
	"github.com/g-celente/AuthService/src/core/repository"
)

type SaveHistoryEvent struct {
	Repo repository.HistoryEventRepository
}

func NewSaveHistoryEvent(repo repository.HistoryEventRepository) *SaveHistoryEvent {
	return &SaveHistoryEvent{Repo: repo}
}

func (s *SaveHistoryEvent) Execute(rec *model.HistoryEvent) error {
	return s.Repo.Save(rec)
}
