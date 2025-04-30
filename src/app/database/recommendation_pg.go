package database

import (
	"github.com/g-celente/AuthService/src/core/model"
	"gorm.io/gorm"
)

type RecommendationPG struct {
	DB *gorm.DB
}

func NewRecommendationPG(db *gorm.DB) *RecommendationPG {
	return &RecommendationPG{DB: db}
}

func (r *RecommendationPG) Save(rec *model.Recommendation) error {
	return r.DB.Create(rec).Error
}
