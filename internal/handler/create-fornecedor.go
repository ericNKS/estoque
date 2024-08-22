package handler

import (
	"net/http"

	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/ericNKS/estoque/internal/service/fornecedor"
	"github.com/gin-gonic/gin"
)

func CreateFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {

	var body types.Fornecedor
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	useCase := fornecedor.CreateFornecedor(repository)
	err := useCase.Execute(body.InstituicaoId, body.NomeFantasia,
		body.RazaoSocial, body.Endereco, body.Cep, body.UnidadeFederativa,
		body.Cnpj, body.InscricaoEstadual, body.Email, body.Telefone)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
	})
}
