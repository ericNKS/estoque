package entities_test

import (
	"testing"

	"github.com/ericNKS/estoque/internal/entities"
)

func TestNewProduto(t *testing.T) {
	fornecedorId := uint64(2)
	instituicaoId := uint64(1)
	nomeProduto := "Tenis de corrida"
	precoCompra := float32(100)
	precoVenda := float32(299.90)
	qtd := uint16(10)
	p, err := entities.NewProduto(
		fornecedorId,
		instituicaoId,
		nomeProduto,
		precoCompra,
		precoVenda,
		qtd,
	)

	if err != nil {
		t.Fatal(err)
	}
	if p.FornecedorId != fornecedorId {
		t.Fatalf("Esperado: %d. Obtido: %d", fornecedorId, p.FornecedorId)
	}
	if p.Nome != nomeProduto {
		t.Fatalf("Esperado: %s. Obtido: %s", nomeProduto, p.Nome)
	}
	if p.PrecoCompra != precoCompra {
		t.Fatalf("Esperado: %f. Obtido: %f", precoCompra, p.PrecoCompra)
	}
	if p.PrecoVenda != precoVenda {
		t.Fatalf("Esperado: %f. Obtido: %f", precoVenda, p.PrecoVenda)
	}
	if p.Quantidade != qtd {
		t.Fatalf("Esperado: %d. Obtido: %d", qtd, p.Quantidade)
	}
}
