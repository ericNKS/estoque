package repository

import (
	"github.com/ericNKS/estoque/internal/db"
	"github.com/ericNKS/estoque/internal/entities/types"
)

func Create(f *types.Fornecedor) error {
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
	return nil, nil
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
