package types

import "gorm.io/gorm"

type Fornecedor struct {
	gorm.Model
	InstituicaoId     uint64 `gorm:"not null" json:"instituicao_id" binding:"required"`
	NomeFantasia      string `gorm:"not null" json:"nome_fantasia" binding:"required"`
	RazaoSocial       string `gorm:"not null" json:"razao_social" binding:"required"`
	Endereco          string `gorm:"not null" json:"endereco" binding:"required"`
	Cep               string `gorm:"not null" json:"cep" binding:"required"`
	UnidadeFederativa string `gorm:"not null" json:"unidade_federativa" binding:"required"`
	Cnpj              string `gorm:"not null" json:"cnpj" binding:"required"`
	InscricaoEstadual string `gorm:"not null" json:"inscricao_estadual" binding:"required"`
	Email             string `json:"email"`
	Telefone          string `json:"telefone"`
}
