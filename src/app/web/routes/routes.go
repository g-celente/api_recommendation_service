package routes

import (
	"github.com/g-celente/AuthService/src/app/web/controller"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/ping", controller.Ping)
	}

}
