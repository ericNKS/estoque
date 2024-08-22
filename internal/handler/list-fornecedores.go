package handler

import (
	"net/http"

	"github.com/ericNKS/estoque/internal/repository"
	"github.com/ericNKS/estoque/internal/service/fornecedor"
	"github.com/gin-gonic/gin"
)

func ListFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	useCase := fornecedor.NewListFornecedor(repository)
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
