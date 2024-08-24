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

func (fr *FornecedorRepository) IsUnique(cnpj string, instituicaoId uint64) (bool, error) {
	db, err := db.Connection()
	if err != nil {
		return false, err
	}
	f := &types.Fornecedor{}
	result := db.First(f, "cnpj = ? AND instituicao_id = ?", cnpj, instituicaoId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return true, nil
		}
		return false, result.Error
	}
	return false, nil
}

func (fr *FornecedorRepository) FindById(id uint64) (*types.Fornecedor, error) {
	db, err := db.Connection()
	if err != nil {
		return nil, err
	}

	f := &types.Fornecedor{}
	result := db.First(f, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return f, nil
}
func (fr *FornecedorRepository) Update(f *types.Fornecedor) error {
	return nil
}
func (fr *FornecedorRepository) Delete(f *types.Fornecedor) error {
	return nil
}
