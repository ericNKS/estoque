package repository

import (
	"github.com/ericNKS/estoque/internal/entities/types"
)

type FornecedorInterface interface {
	Create(f *types.Fornecedor) error
	FindAll() ([]*types.Fornecedor, error)
	FindById(id uint) (*types.Fornecedor, error)
	Update(f *types.Fornecedor) error
	Delete(f *types.Fornecedor) error
}
