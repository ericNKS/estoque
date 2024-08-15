package repository

import (
	"github.com/ericNKS/estoque/internal/db"
	"github.com/ericNKS/estoque/internal/entities/types"
	"gorm.io/gorm"
)

type FornecedorRepository struct {
	db *gorm.DB
}

func NewFornecedorRepository() (*FornecedorRepository, error) {
	db, err := db.Connection()
	if err != nil {
		return nil, err
	}

	return &FornecedorRepository{
		db: db,
	}, nil
}

func (fr *FornecedorRepository) Create(f *types.Fornecedor) error {
	result := fr.db.Create(f)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (fr *FornecedorRepository) FindAll() ([]*types.Fornecedor, error) {
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

func (fr *FornecedorRepository) FindById(id uint) (*types.Fornecedor, error) {
	return nil, nil
}
func (fr *FornecedorRepository) Update(f *types.Fornecedor) error {
	return nil
}
func (fr *FornecedorRepository) Delete(f *types.Fornecedor) error {
	return nil
}
