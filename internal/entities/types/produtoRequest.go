package types

type CreateProdutoRequest struct {
	InstituicaoId uint64 `json:"instituicao_id" binding:"required"`
	FornecedorId  string `json:"nome_fantasia" binding:"required"`
	CodigoBarras  string `json:"endereco" binding:"required"`
	Nome          string `json:"cep" binding:"required"`
	PrecoCompra   string `json:"unidade_federativa" binding:"required"`
	PrecoVenda    string `json:"cnpj" binding:"required"`
	Quantidade    string `json:"inscricao_estadual" binding:"required"`
}
