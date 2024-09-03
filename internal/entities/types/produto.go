package types

import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	FornecedorId uint64     `gorm:"not null" json:"fornecedor_id" binding:"required"`
	Fornecedor   Fornecedor `gorm:"references:FornecedorId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Nome         string     `gorm:"not null" json:"nome" binding:"required"`
	Preco        float64    `gorm:"not null" json:"preco" binding:"required" default:"0"`
	Quantidade   uint64     `json:"quantidade" binding:"required" default:"0"`
}
