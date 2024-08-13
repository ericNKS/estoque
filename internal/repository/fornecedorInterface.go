package repository

import "estoque/internal/entities"

type FornecedorInterface interface {
	Create(f *entities.Fornecedor) error
	FindAll() ([]*entities.Fornecedor, error)
	FindById(id string) (entities.Fornecedor, error)
	Update(f *entities.Fornecedor) error
	Delete(f *entities.Fornecedor) error
}
