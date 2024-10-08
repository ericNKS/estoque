package entities

import (
	"fmt"

	"github.com/ericNKS/estoque/internal/entities/types"
)

func NewFornecedor(instituicaoId uint64, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual, email, telefone string) (*types.Fornecedor, error) {

	err := validateFornecedor(instituicaoId, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual)
	if err != nil {
		return nil, err
	}

	return &types.Fornecedor{
		InstituicaoId:     instituicaoId,
		NomeFantasia:      nomeFantasia,
		RazaoSocial:       razaoSocial,
		Endereco:          endereco,
		Cep:               cep,
		UnidadeFederativa: unidadeFederativa,
		Cnpj:              cnpj,
		InscricaoEstadual: inscricaoEstadual,
		Email:             email,
		Telefone:          telefone,
	}, nil
}

func validateFornecedor(instituicaoId uint64, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual string) error {
	if nomeFantasia == "" {
		return fmt.Errorf("nome fantasia é obrigatorio")
	}

	if instituicaoId == 0 {
		return fmt.Errorf("instituicao id é obrigatorio")
	}

	if razaoSocial == "" {
		return fmt.Errorf("razao Social é obrigatorio")
	}

	if endereco == "" {
		return fmt.Errorf("endereco é obrigatorio")
	}

	if cep == "" {
		return fmt.Errorf("CEP é obrigatorio")
	}

	if unidadeFederativa == "" {
		return fmt.Errorf("unidade federativa é obrigatorio")
	}

	if cnpj == "" {
		return fmt.Errorf("CNPJ é obrigatorio")
	}

	if inscricaoEstadual == "" {
		return fmt.Errorf("inscricao estadual é obrigatorio")
	}
	return nil
}
