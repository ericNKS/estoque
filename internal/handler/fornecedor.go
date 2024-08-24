package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ericNKS/estoque/internal/repository"
	"github.com/ericNKS/estoque/internal/service/fornecedor"
	"github.com/gin-gonic/gin"
)

func CreateFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	var body CreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	useCase := fornecedor.CreateFornecedor(repository)
	err := useCase.Execute(
		body.InstituicaoId,
		body.NomeFantasia,
		body.RazaoSocial,
		body.Endereco,
		body.Cep,
		body.UnidadeFederativa,
		body.Cnpj,
		body.InscricaoEstadual,
		body.Email,
		body.Telefone,
	)

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

func DeleteFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	teste := ctx.Params
	fmt.Println(teste)
}

func ListFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	useCase := fornecedor.ListFornecedor(repository)
	fornecedores, err := useCase.Execute()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": fornecedores,
	})
}

func UpdateFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var request UpdateRequest
	if err = ctx.ShouldBindBodyWithJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	findCase := fornecedor.FindById(repository)

	f, err := findCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": f})
}
