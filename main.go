package main

import (
	"github.com/g-celente/AuthService/src/app/config"
	"github.com/g-celente/AuthService/src/app/database"
	"github.com/g-celente/AuthService/src/app/web/controller"
	"github.com/g-celente/AuthService/src/app/web/routes"
	"github.com/g-celente/AuthService/src/core/useCase"
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

	db, err := config.Init()

	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	server := gin.Default()

	repo := database.NewRecommendationPG(db)
	use := useCase.NewSaveRecommendation(repo)
	controller := controller.NewRecommendationController(use)

	routes.RegisterRoutes(server, controller)

	server.Run(":8000")

}
