package handler

import (
	"fmt"

	"github.com/ericNKS/estoque/internal/repository"
	"github.com/gin-gonic/gin"
)

func DeleteFornecedor(ctx *gin.Context, repository *repository.FornecedorRepository) {
	teste := ctx.Params
	fmt.Println(teste)
}
