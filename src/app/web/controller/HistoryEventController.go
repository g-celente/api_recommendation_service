package controller

import (
	"net/http"

	"github.com/g-celente/AuthService/src/core/model"
	"github.com/g-celente/AuthService/src/core/useCase"
	"github.com/gin-gonic/gin"
)

type HistoryEventController struct {
	UseCase *useCase.SaveHistoryEvent
}

func NewHistoryEventController(use *useCase.SaveHistoryEvent) *HistoryEventController {
	return &HistoryEventController{UseCase: use}
}

func (ctrl *HistoryEventController) Receive(c *gin.Context) {
	var input model.HistoryEvent

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
