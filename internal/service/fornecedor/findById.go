package fornecedor

import (
	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/ericNKS/estoque/internal/repository"
)

type findByIdFornecedor struct {
	repo repository.FornecedorInterface
}

func FindFornecedor(repo repository.FornecedorInterface) *findByIdFornecedor {
	return &findByIdFornecedor{
		repo: repo,
	}
}

func (u *findByIdFornecedor) Execute(id uint64) (*types.Fornecedor, error) {

	f, err := u.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return f, nil
}
