package fornecedor

import (
	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/ericNKS/estoque/internal/repository"
)

type deleteFornecedor struct {
	repo repository.FornecedorInterface
}

func DeleteFornecedor(repo repository.FornecedorInterface) *deleteFornecedor {
	return &deleteFornecedor{
		repo: repo,
	}
}

func (d *deleteFornecedor) Validate(id uint64) (*types.Fornecedor, error) {
	f, err := d.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (d *deleteFornecedor) Execute(f *types.Fornecedor) error {
	err := d.repo.Delete(f)
	if err != nil {
		return err
	}

	return nil
}
