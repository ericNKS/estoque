package entities

import (
	"fmt"

	"github.com/ericNKS/estoque/internal/entities/types"
)

func NewProduto(fornecedorId uint64, nome string, preco float32, quantidade uint16) (*types.Produto, error) {
	err := validateProduto(fornecedorId, nome, preco, quantidade)
	if err != nil {
		return nil, err
	}
	p := types.Produto{
		FornecedorId: fornecedorId,
		Nome:         nome,
		Preco:        preco,
		Quantidade:   quantidade,
	}
	return &p, nil
}
func validateProduto(fornecedorId uint64, nome string, preco float32, quantidade uint16) error {
	if fornecedorId == 0 {
		return fmt.Errorf("fornecedorId invalido")
	}
	if nome == "" {
		return fmt.Errorf("nome invalido")
	}
	if preco < 0 {
		return fmt.Errorf("preco invalido")
	}
	if quantidade < 0 {
		return fmt.Errorf("quantidade invalida")
	}

	return nil
}
