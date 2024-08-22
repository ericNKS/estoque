package fornecedor

import (
	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/ericNKS/estoque/internal/repository"
)

type listFornecedor struct {
	repo repository.FornecedorInterface
}

func ListFornecedor(repo repository.FornecedorInterface) *listFornecedor {
	return &listFornecedor{
		repo: repo,
	}
}

func (u *listFornecedor) Execute() ([]*types.Fornecedor, error) {

	fornecedores, err := u.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return fornecedores, nil
}
