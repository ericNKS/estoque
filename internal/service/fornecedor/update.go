package fornecedor

import (
	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/ericNKS/estoque/internal/repository"
)

type updateFornecedor struct {
	repo repository.FornecedorInterface
}

func UpdateFornecedor(repo repository.FornecedorInterface) *updateFornecedor {
	return &updateFornecedor{
		repo: repo,
	}
}

func (u *updateFornecedor) Execute(f *types.Fornecedor) (*types.Fornecedor, error) {
	err := u.repo.Update(f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
