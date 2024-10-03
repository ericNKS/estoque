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

func (u *updateFornecedor) Validate(id uint64) (*types.Fornecedor, error) {
	f, err := u.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (u *updateFornecedor) Execute(f *types.Fornecedor, ur *types.UpdateFornecedorRequest) (*types.Fornecedor, error) {

	if ur.Cep != "" {
		f.Cep = ur.Cep
	}

	if ur.Cnpj != "" {
		f.Cnpj = ur.Cnpj
	}

	if ur.Email != "" {
		f.Email = ur.Email
	}

	if ur.Endereco != "" {
		f.Endereco = ur.Endereco
	}

	if ur.InscricaoEstadual != "" {
		f.InscricaoEstadual = ur.InscricaoEstadual
	}

	if ur.NomeFantasia != "" {
		f.NomeFantasia = ur.NomeFantasia
	}

	if ur.RazaoSocial != "" {
		f.RazaoSocial = ur.RazaoSocial
	}

	if ur.Telefone != "" {
		f.Telefone = ur.Telefone
	}

	if ur.UnidadeFederativa != "" {
		f.UnidadeFederativa = ur.UnidadeFederativa
	}

	err := u.repo.Update(f)
	if err != nil {
		return nil, err
	}
	return f, nil
}
