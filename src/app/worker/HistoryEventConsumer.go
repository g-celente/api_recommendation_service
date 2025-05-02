package worker

import (
	"encoding/json"
	"log"

	"github.com/g-celente/AuthService/src/app/config/messagebroker"
	"github.com/g-celente/AuthService/src/core/model"
	"github.com/g-celente/AuthService/src/core/useCase"
)

func ConsumeHistoryEvents(broker *messagebroker.RabbitMQ, usecase *useCase.SaveHistoryEvent) {
	msgs, err := broker.ConsumeMessages()
	if err != nil {
		log.Fatalf("Erro ao consumir mensagens: %v", err)
	}

	go func() {
		for msg := range msgs {
			var event model.HistoryEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Printf("Erro ao converter mensagem para modelo: %v", err)
				continue
			}

			if err := usecase.Execute(&event); err != nil {
				log.Printf("Erro ao salvar evento: %v", err)
				continue
			}

			log.Printf("Evento salvo com sucesso: %+v", event)
		}
	}()
}
