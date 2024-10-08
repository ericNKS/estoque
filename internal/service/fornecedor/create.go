package fornecedor

import (
	"github.com/ericNKS/estoque/internal/entities"
	"github.com/ericNKS/estoque/internal/repository"
)

type createFornecedor struct {
	repo repository.FornecedorInterface
}

func CreateFornecedor(repo repository.FornecedorInterface) *createFornecedor {
	return &createFornecedor{
		repo: repo,
	}
}

func (u *createFornecedor) Execute(instituicaoId uint64, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual, email, telefone string) error {

	f, err := entities.NewFornecedor(instituicaoId, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual, email, telefone)
	if err != nil {
		return err
	}

	err = u.repo.IsUnique(f.Cnpj, f.InstituicaoId)
	if err != nil {
		return err
	}

	if err = u.repo.Create(f); err != nil {
		return err
	}
	return nil
}
