package repository

import (
	"estoque/internal/entities"
	"testing"
)

func TestCreateFornecedor(t *testing.T) {
	f, err := entities.NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}
	if err = Create(f); err != nil {
		t.Fatal(err)
	}
}

func TestFindAllFornecedor(t *testing.T) {
	_, err := FindAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestFindByIdFornecedor(t *testing.T) {
	f, err := entities.NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}

	uF, err := FindById(f.ID)
	if err != nil {
		t.Fatal(err)
	}
	if uF.ID != f.ID {
		t.Fatal("Fornecedores diferentes")
	}
}

func TestUpdateFornecedor(t *testing.T) {
	f, err := entities.NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}
	if err = Create(f); err != nil {
		t.Fatal(err)
	}

	f.Email = "gerbete_novo@gmail.com"
	if err = Update(f); err != nil {
		t.Fatal(err)
	}
}

func TestDeleteFornecedor(t *testing.T) {
	f, err := entities.NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}
	if err = Create(f); err != nil {
		t.Fatal(err)
	}

	if err = Delete(f.ID); err != nil {
		t.Fatal(err)
	}

	nF, err := FindById(f.ID)
	if err != nil {
		t.Fatal(err)
	}
	if nF.ID == f.ID {
		t.Fatal("Fornecedores nao foi deletado")
	}
}
