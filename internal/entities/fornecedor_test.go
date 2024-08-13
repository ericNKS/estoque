package entities

import "testing"

func TestNewFornecedor(t *testing.T) {
	// Nome fantasia - razao social - endereco - cep - uf - cnpj - inscricao estadual - email - telefone
	// email e telefone podem ser vazios mas se nao forem devem ser verificados sua veracidade

	_, err := NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "71992037328")
	if err != nil {
		t.Fatal(err)
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "", "71992037328")
	if err == nil {
		t.Fatal("Numero de telefone pode ser vazio")
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "gerbete@gmail.com", "")
	if err == nil {
		t.Fatal("Numero de telefone pode ser vazio")
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "", "", "")
	if err == nil {
		t.Fatal("A inscricao estadual não pode ser vazio")
	}
	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "", "1234567890", "", "")
	if err == nil {
		t.Fatal("O CNPJ não pode ser vazio")
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "", "12345678901234", "1234567890", "", "")
	if err == nil {
		t.Fatal("A unidade federativa não pode ser vazio")
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "", "SP", "12345678901234", "1234567890", "", "")
	if err == nil {
		t.Fatal("O CEP não pode ser vazio")
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "", "12345678", "SP", "12345678901234", "1234567890", "", "")
	if err == nil {
		t.Fatal("O endereco não pode ser vazio")
	}

	_, err = NewFornecedor("", "", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "", "")
	if err == nil {
		t.Fatal("A razão social não pode ser vazio")
	}

	_, err = NewFornecedor("Gerbete Atacado", "Gerbete", "Rua Jayme Gonçalves", "12345678", "SP", "12345678901234", "1234567890", "", "")
	if err == nil {
		t.Fatal("Nome fantasia não pode ser vazio")
	}

}
