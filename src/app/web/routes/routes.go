package routes

import (
	"github.com/g-celente/AuthService/src/app/web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, controller *controller.HistoryEventController) {

	api := r.Group("/api")
	{
		api.POST("/recommendation", controller.Receive)
	}

}
