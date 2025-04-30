package repository

import "github.com/g-celente/AuthService/src/core/model"

type RecommendationRepository interface {
	Save(recommendation *model.Recommendation) error
}
