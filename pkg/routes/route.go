package routes

import (
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/ericNKS/estoque/pkg/handler"
	"github.com/gin-gonic/gin"
)

func FornecedorRoute(r *gin.Engine) {
	fr, err := repository.NewFornecedorRepository()
	if err != nil {
		panic(err)
	}
	fornecedorGroup := r.Group("/fornecedor")

	fornecedorGroup.POST("/", func(c *gin.Context) {
		handler.CreateFornecedor(c, fr)
	})

	fornecedorGroup.GET("/", func(c *gin.Context) {
		handler.ListFornecedor(c, fr)
	})
}
