package types

import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	CodigoBarras string     `gorm:"unique;not null" json:"codigo_barras" binding:"required"`
	FornecedorId uint64     `gorm:"not null" json:"fornecedor_id" binding:"required"`
	Fornecedor   Fornecedor `gorm:"references:FornecedorId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Nome         string     `gorm:"not null" json:"nome" binding:"required"`
	Preco        float32    `gorm:"not null" json:"preco" binding:"required" default:"0"`
	Quantidade   uint16     `json:"quantidade" binding:"required" default:"0"`
}
