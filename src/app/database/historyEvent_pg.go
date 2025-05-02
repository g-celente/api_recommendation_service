package database

import (
	"github.com/g-celente/AuthService/src/core/model"
	"gorm.io/gorm"
)

// HistoryEventPG é o repositório que interage com o banco de dados
type HistoryEventPG struct {
	DB *gorm.DB
}

// NewHistoryEventPG cria uma nova instância do repositório HistoryEventPG
func NewHistoryEventPG(db *gorm.DB) *HistoryEventPG {
	return &HistoryEventPG{DB: db}
}

// Save salva um evento histórico no banco de dados
func (r *HistoryEventPG) Save(rec *model.HistoryEvent) error {
	return r.DB.Create(rec).Error
}

// GetTopThemesByEmail busca os principais temas associados ao e-mail do assinante
func (r *HistoryEventPG) GetTopThemesByEmail(email string, limit int) ([]string, error) {
	// Define um slice para armazenar os temas
	var themes []string

	// Consulta no banco de dados para pegar os temas associados ao email
	err := r.DB.Table("history_events").
		Select("event_theme").
		Where("email_subscriber = ?", email).
		Limit(limit).
		Pluck("event_theme", &themes).Error

	// Retorna os temas ou o erro
	return themes, err
}
