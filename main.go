package main

import (
	"github.com/g-celente/AuthService/src/app/web/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")

}
