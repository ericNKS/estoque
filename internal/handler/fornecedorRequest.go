package handler

type UpdateRequest struct {
	InstituicaoId     uint64 `json:"instituicao_id" binding:"required"`
	NomeFantasia      string `json:"nome_fantasia"`
	RazaoSocial       string `json:"razao_social"`
	Endereco          string `json:"endereco"`
	Cep               string `json:"cep"`
	UnidadeFederativa string `json:"unidade_federativa"`
	Cnpj              string `json:"cnpj"`
	InscricaoEstadual string `json:"inscricao_estadual"`
	Email             string `json:"email"`
	Telefone          string `json:"telefone"`
}
type CreateRequest struct {
	InstituicaoId     uint64 `json:"instituicao_id" binding:"required"`
	NomeFantasia      string `json:"nome_fantasia" binding:"required"`
	RazaoSocial       string `json:"razao_social" binding:"required"`
	Endereco          string `json:"endereco" binding:"required"`
	Cep               string `json:"cep" binding:"required"`
	UnidadeFederativa string `json:"unidade_federativa" binding:"required"`
	Cnpj              string `json:"cnpj" binding:"required"`
	InscricaoEstadual string `json:"inscricao_estadual" binding:"required"`
	Email             string `json:"email"`
	Telefone          string `json:"telefone"`
}
