package controller

import (
	"net/http"

	"github.com/g-celente/AuthService/src/core/useCase"
	"github.com/gin-gonic/gin"
)

type RecommendationController struct {
	UseCase *useCase.RecommendationUseCase
}

func NewRecommendationController(use *useCase.RecommendationUseCase) *RecommendationController {
	return &RecommendationController{UseCase: use}
}

func (ctrl *RecommendationController) Recommend(c *gin.Context) {
	email := c.Param("email")

	events, err := ctrl.UseCase.Execute(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, events)
}
