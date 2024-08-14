package types

import "gorm.io/gorm"

type Fornecedor struct {
	gorm.Model
	InstituicaoId     uint64 `gorm:"not null" json:"instituicao_id"`
	NomeFantasia      string `gorm:"not null" json:"nome_fantasia"`
	RazaoSocial       string `gorm:"not null" json:"razao_social"`
	Endereco          string `gorm:"not null" json:"endereco"`
	Cep               string `gorm:"not null" json:"cep"`
	UnidadeFederativa string `gorm:"not null" json:"unidade_federativa"`
	Cnpj              string `gorm:"not null" json:"cnpj"`
	InscricaoEstadual string `gorm:"not null" json:"inscricao_estadual"`
	Email             string `json:"email"`
	Telefone          string `json:"telefone"`
}
