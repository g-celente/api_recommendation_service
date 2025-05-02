package routes

import (
	"github.com/g-celente/AuthService/src/app/web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, recController *controller.RecommendationController) {
	api := r.Group("/api")
	{
		api.GET("/recommendation/:email", recController.Recommend)
	}
}
