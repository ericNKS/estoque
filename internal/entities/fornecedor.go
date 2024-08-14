package entities

import (
	"fmt"

	"gorm.io/gorm"
)

// Nome fantasia - razao social - endereco - UnidadeFederativa - uf - cnpj - inscricao estadual - email - telefone
type Fornecedor struct {
	gorm.Model
	InstituicaoId     uint64 `gorm:"not null" json:"instituicao_id"`
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

func NewFornecedor(instituicaoId uint64, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual, email, telefone string) (*Fornecedor, error) {

	err := validateFornecedor(instituicaoId, nomeFantasia, razaoSocial, endereco, cep, unidadeFederativa, cnpj, inscricaoEstadual)
	if err != nil {
		return nil, err
	}

	return &Fornecedor{
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
