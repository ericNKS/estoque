package entities

import (
	"fmt"

	"gorm.io/gorm"
)

// Nome fantasia - razao social - endereco - UnidadeFederativa - uf - cnpj - inscricao estadual - email - telefone
type fornecedor struct {
	gorm.Model
	NomeFantasia      string `gorm:"not null" json:"nome_fantasia"`
	RazaoSocial       string `gorm:"not null" json:"razao_social"`
	Endereco          string `gorm:"not null" json:"endereco"`
	Cep               string `gorm:"not null" json:"cep"`
	UnidadeFederativa string `gorm:"not null" json:"unidade_federativa"`
	Cnpj              string `gorm:"not null; unique" json:"cnpj"`
	InscricaoEstadual string `gorm:"not null; unique" json:"inscricao_estadual"`
	Email             string `json:"email"`
	Telefone          string `json:"telefone"`
}

func NewFornecedor(nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual, email, telefone string) (*fornecedor, error) {

	err := validateFornecedor(nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual)
	if err != nil {
		return nil, err
	}

	return &fornecedor{
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

func validateFornecedor(nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual string) error {
	if nomeFantasia == "" {
		return fmt.Errorf("nome fantasia é obrigatorio")
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
