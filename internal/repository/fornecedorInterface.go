package repository

import (
	"github.com/ericNKS/estoque/internal/entities/types"
)

type FornecedorInterface interface {
	Create(f *types.Fornecedor) error
	FindAll() ([]*types.Fornecedor, error)
	FindById(id uint64) (*types.Fornecedor, error)
	IsUnique(cnpj string, instituicaoId uint64) error
	Update(f *types.Fornecedor) error
	Delete(f *types.Fornecedor) error
}
