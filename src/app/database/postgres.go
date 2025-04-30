package database

import (
	"fmt"
	"os"

	"github.com/g-celente/AuthService/src/core/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar no Postgres: %v", err)
	}

	err = db.AutoMigrate(&model.Recommendation{})

	if err != nil {
		return nil, fmt.Errorf("falha ao fazer migration: %v ", err)
	}

	return db, nil
}
