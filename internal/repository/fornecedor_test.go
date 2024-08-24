package repository

import (
	"fmt"
	"testing"

	"github.com/ericNKS/estoque/internal/entities"
	"github.com/joho/godotenv"
)

func TestCreateFornecedor(t *testing.T) {
	godotenv.Load("../../.env")
	f, err := entities.NewFornecedor(3, "Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}

	fr, err := NewFornecedorRepository()
	if err != nil {
		t.Fatal(err)
	}

	if err = fr.Create(f); err != nil {
		t.Fatal(err)
	}
	fmt.Println(f)
}

func TestFindAllFornecedor(t *testing.T) {
	godotenv.Load("../../.env")
	fr, err := NewFornecedorRepository()
	if err != nil {
		t.Fatal(err)
	}
	_, err = fr.FindAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindByIdFornecedor(t *testing.T) {
	godotenv.Load("../../.env")
	f, err := entities.NewFornecedor(3, "Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}

	fr, err := NewFornecedorRepository()
	if err != nil {
		t.Fatal(err)
	}

	uF, err := fr.FindById(uint64(f.ID))
	if err != nil {
		t.Fatal(err)
	}
	if uF.ID != f.ID {
		t.Fatal("Fornecedores diferentes")
	}
}

func TestUpdateFornecedor(t *testing.T) {
	godotenv.Load("../../.env")
	f, err := entities.NewFornecedor(3, "Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}
	fr, err := NewFornecedorRepository()
	if err != nil {
		t.Fatal(err)
	}
	if err = fr.Create(f); err != nil {
		t.Fatal(err)
	}

	f.Email = "gerbete_novo@gmail.com"
	if err = fr.Update(f); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteFornecedor(t *testing.T) {
	godotenv.Load("../../.env")
	fr, err := NewFornecedorRepository()
	if err != nil {
		t.Fatal(err)
	}
	f, err := entities.NewFornecedor(3, "Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}
	if err = fr.Create(f); err != nil {
		t.Fatal(err)
	}

	if err = fr.Delete(f); err != nil {
		t.Fatal(err)
	}

	nF, err := fr.FindById(uint64(f.ID))
	if err != nil {
		t.Fatal(err)
	}
	if nF.ID == f.ID {
		t.Fatal("Fornecedores nao foi deletado")
	}
}
