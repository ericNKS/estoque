package types

import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	CodigoBarras  string     `gorm:"unique;not null" json:"codigo_barras" binding:"required"`
	FornecedorId  uint64     `gorm:"not null" json:"fornecedor_id" binding:"required"`
	Fornecedor    Fornecedor `gorm:"references:FornecedorId,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	InstituicaoId uint64     `gorm:"not null" json:"instituicao_id" binding:"required"`
	Nome          string     `gorm:"not null" json:"nome" binding:"required"`
	PrecoCompra   float32    `gorm:"not null" json:"preco_compra" binding:"required" default:"0"`
	PrecoVenda    float32    `gorm:"not null" json:"preco_venda" binding:"required" default:"0"`
	Quantidade    uint16     `json:"quantidade" binding:"required" default:"0"`
}
