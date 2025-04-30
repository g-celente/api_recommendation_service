package useCase

import (
	"github.com/g-celente/AuthService/src/core/model"
	"github.com/g-celente/AuthService/src/core/repository"
)

type SaveRecommendation struct {
	Repo repository.RecommendationRepository
}

func NewSaveRecommendation(repo repository.RecommendationRepository) *SaveRecommendation {
	return &SaveRecommendation{Repo: repo}
}

func (s *SaveRecommendation) Execute(rec *model.Recommendation) error {
	return s.Repo.Save(rec)
}
