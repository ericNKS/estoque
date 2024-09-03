package types

import "gorm.io/gorm"

type Produto struct {
	gorm.Model
	FornecedorId uint64
	Fornecedor   Fornecedor `gorm:"references:FornecedorId"`
	Nome         string
	Preco        float64
	Quantidade   uint64
}
