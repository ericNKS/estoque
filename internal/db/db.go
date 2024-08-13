package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connection() (*sqlx.DB, error) {
	godotenv.Load("../../.env")
	host := os.Getenv("POSTGRES_HOST")
	dbName := os.Getenv("POSTGRES_DB_NAME")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dns := fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s", user, dbName, password, host)

	db, err := sqlx.Connect("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}
