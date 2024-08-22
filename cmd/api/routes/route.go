package routes

import (
	"github.com/ericNKS/estoque/internal/handler"
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/gin-gonic/gin"
)

func FornecedorRoute(r *gin.Engine) {
	fr, err := repository.NewFornecedorRepository()
	if err != nil {
		panic(err)
	}
	routeApiV1 := r.Group("/api/v1")
	fornecedorGroup := routeApiV1.Group("/fornecedor")

	fornecedorGroup.POST("/", func(c *gin.Context) {
		handler.CreateFornecedor(c, fr)
	})

	fornecedorGroup.GET("/", func(c *gin.Context) {
		handler.ListFornecedor(c, fr)
	})

	fornecedorGroup.DELETE("/:id", func(c *gin.Context) {
		handler.DeleteFornecedor(c, fr)
	})

}
