package controller

import (
	"net/http"

	"github.com/g-celente/AuthService/src/core/model"
	"github.com/g-celente/AuthService/src/core/useCase"
	"github.com/gin-gonic/gin"
)

type RecommendationController struct {
	UseCase *useCase.SaveRecommendation
}

func NewRecommendationController(use *useCase.SaveRecommendation) *RecommendationController {
	return &RecommendationController{UseCase: use}
}

func (ctrl *RecommendationController) Receive(c *gin.Context) {
	var input model.Recommendation

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.UseCase.Execute(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recomendação salva"})
}
