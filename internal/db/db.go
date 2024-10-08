package db

import (
	"fmt"
	"os"

	"github.com/ericNKS/estoque/internal/entities/types"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dns := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", user, dbName, password, host)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&types.Fornecedor{})

	return db, nil
}
