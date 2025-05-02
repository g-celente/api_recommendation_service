package main

import (
	"os"

	"github.com/g-celente/AuthService/src/app/config"
	"github.com/g-celente/AuthService/src/app/config/messagebroker"
	"github.com/g-celente/AuthService/src/app/database"
	"github.com/g-celente/AuthService/src/app/web/controller"
	"github.com/g-celente/AuthService/src/app/web/routes"
	"github.com/g-celente/AuthService/src/app/worker"
	"github.com/g-celente/AuthService/src/core/useCase"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger := config.GetLogger("main")

	if err := godotenv.Load(); err != nil {
		logger.Errorf("Erro ao carregar .env: %v", err)
		return
	}

	db, err := config.Init()
	if err != nil {
		logger.Errorf("Erro ao inicializar banco: %v", err)
		return
	}

	repo := database.NewHistoryEventPG(db)
	saveUse := useCase.NewSaveHistoryEvent(repo)
	recUse := useCase.NewRecommendationUseCase(repo)

	// Inicia RabbitMQ
	rabbitmq, err := messagebroker.NewRabbitMQ(os.Getenv("RABBITMQ_CONNECTION_URL"), os.Getenv("RABBITMQ_QUEUE_NAME"))
	if err != nil {
		logger.Errorf("Erro ao conectar no RabbitMQ: %v", err)
		return
	}
	defer rabbitmq.Close()

	// Consome mensagens em background
	worker.ConsumeHistoryEvents(rabbitmq, saveUse)

	// Inicia API
	server := gin.Default()
	recCtrl := controller.NewRecommendationController(recUse)

	routes.RegisterRoutes(server, recCtrl)
	server.Run(":8000")
}
