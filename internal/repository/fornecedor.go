package repository

import (
	"github.com/ericNKS/estoque/internal/db"
	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/joho/godotenv"
)

func Create(f *types.Fornecedor) error {
	godotenv.Load("../../.env")
	db, err := db.Connection()
	if err != nil {
		return err
	}
	result := db.Create(f)
	if result.Error != nil {
		return err
	}

	return nil
}

func FindAll() ([]*types.Fornecedor, error) {
	godotenv.Load("../../.env")
	db, err := db.Connection()
	if err != nil {
		return nil, err
	}

	var fornecedores []*types.Fornecedor
	result := db.Find(&fornecedores)
	if result.Error != nil {
		return nil, err
	}

	return fornecedores, nil
}

func FindById(id uint) (*types.Fornecedor, error) {
	return nil, nil
}
func Update(f *types.Fornecedor) error {
	return nil
}
func Delete(f *types.Fornecedor) error {
	return nil
}
