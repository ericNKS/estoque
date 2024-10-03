package entities

import (
	"fmt"

	"github.com/ericNKS/estoque/internal/entities/types"
)

func NewProduto(fornecedorId, instituicaoId uint64, nome string, precoCompra, PrecoVenda float32, quantidade uint16) (*types.Produto, error) {
	err := validateProduto(fornecedorId, instituicaoId, nome, precoCompra, PrecoVenda, quantidade)
	if err != nil {
		return nil, err
	}
	p := types.Produto{
		FornecedorId: fornecedorId,
		Nome:         nome,
		PrecoCompra:  precoCompra,
		PrecoVenda:   PrecoVenda,
		Quantidade:   quantidade,
	}
	return &p, nil
}
func validateProduto(fornecedorId, instituicaoId uint64, nome string, precoCompra, precoVenda float32, quantidade uint16) error {
	if fornecedorId <= 0 {
		return fmt.Errorf("fornecedorId invalido")
	}
	if instituicaoId <= 0 {
		return fmt.Errorf("fornecedorId invalido")
	}
	if nome == "" {
		return fmt.Errorf("nome invalido")
	}
	if precoCompra < 0 {
		return fmt.Errorf("preco invalido")
	}
	if precoVenda < 0 {
		return fmt.Errorf("preco invalido")
	}
	if quantidade < 0 {
		return fmt.Errorf("quantidade invalida")
	}

	return nil
}
