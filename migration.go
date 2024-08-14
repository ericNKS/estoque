package main

import (
	"github.com/ericNKS/estoque/internal/db"
	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	db, err := db.Connection()

	if err != nil {
		panic(err)
	}

	if !db.Migrator().HasTable(&types.Fornecedor{}) {
		err = db.AutoMigrate(&types.Fornecedor{})
		if err != nil {
			panic(err)
		}
	}
}
