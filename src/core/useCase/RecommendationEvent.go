package useCase

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/g-celente/AuthService/src/core/model"
	"github.com/g-celente/AuthService/src/core/repository"
)

type RecommendationUseCase struct {
	Repo repository.HistoryEventRepository
}

func NewRecommendationUseCase(repo repository.HistoryEventRepository) *RecommendationUseCase {
	return &RecommendationUseCase{Repo: repo}
}

func (r *RecommendationUseCase) Execute(email string) ([]model.EventRecommendation, error) {
	if email == "" {
		return nil, errors.New("e-mail não pode estar vazio")
	}

	themes, err := r.Repo.GetTopThemesByEmail(email, 3)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar temas para o e-mail informado: %w", err)
	}
	if len(themes) == 0 {
		return nil, fmt.Errorf("nenhum tema encontrado para o e-mail informado")
	}

	query := strings.Join(themes, ",")
	urlBase := os.Getenv("EVENT_SERVICE_URL")
	if urlBase == "" {
		return nil, errors.New("variável de ambiente EVENT_SERVICE_URL não configurada")
	}

	url := fmt.Sprintf("%s?themes=%s", urlBase, query)
	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar eventos do serviço externo: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("serviço de eventos retornou status %d", resp.StatusCode)
	}

	var eventos []model.EventRecommendation
	if err := json.NewDecoder(resp.Body).Decode(&eventos); err != nil {
		return nil, fmt.Errorf("erro ao decodificar resposta do serviço de eventos: %w", err)
	}

	return eventos, nil
}
