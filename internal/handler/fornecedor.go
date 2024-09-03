package handler

import (
	"net/http"
	"strconv"

	"github.com/ericNKS/estoque/internal/entities/types"
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/ericNKS/estoque/internal/service/fornecedor"
	"github.com/gin-gonic/gin"
)

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

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func FindFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	useCase := fornecedor.FindFornecedor(repository)
	f, err := useCase.Execute(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": f,
	})
}

func CreateFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	var body types.CreateRequest
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
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := fornecedor.DeleteFornecedor(repository)
	f, err := useCase.Validate(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = useCase.Execute(f)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

func UpdateFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var r types.UpdateRequest
	if err = ctx.ShouldBindBodyWithJSON(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	useCase := fornecedor.UpdateFornecedor(repository)
	f, err := useCase.Validate(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	f, err = useCase.Execute(f, &r)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"success": f})
}
