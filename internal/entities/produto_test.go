package entities_test

import (
	"testing"

	"github.com/ericNKS/estoque/internal/entities"
)

func TestNewProduto(t *testing.T) {
	fornecedorId := uint64(2)
	nomeProduto := "Tenis de corrida"
	preco := float32(299.90)
	qtd := uint16(10)
	p, err := entities.NewProduto(
		fornecedorId,
		nomeProduto,
		preco,
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
	if p.Preco != preco {
		t.Fatalf("Esperado: %f. Obtido: %f", preco, p.Preco)
	}
	if p.Quantidade != qtd {
		t.Fatalf("Esperado: %d. Obtido: %d", qtd, p.Quantidade)
	}
}
