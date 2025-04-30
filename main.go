package main

import (
	"github.com/g-celente/AuthService/src/app/config"
	"github.com/g-celente/AuthService/src/app/web/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	errEnv := godotenv.Load()
	if errEnv != nil {
		logger.Errorf("Erro ao carregar variáveis de ambiente: %v", errEnv)
		return
	}

	err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8000")

}
