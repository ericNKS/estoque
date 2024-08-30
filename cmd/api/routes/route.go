package routes

import (
	"github.com/ericNKS/estoque/internal/handler"
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/gin-gonic/gin"
)

func fRepository() *repository.FornecedorRepository {
	fr, err := repository.NewFornecedorRepository()
	if err != nil {
		panic(err)
	}
	return fr
}
func FornecedorRoute(r *gin.Engine) {
	routeApiV1 := r.Group("/api/v1")
	fornecedorGroup := routeApiV1.Group("/fornecedor")

	fornecedorGroup.POST("/", func(c *gin.Context) {
		handler.CreateFornecedor(c, fRepository())
	})

	fornecedorGroup.GET("/", func(c *gin.Context) {
		handler.ListFornecedor(c, fRepository())
	})

	fornecedorGroup.DELETE("/:id", func(c *gin.Context) {
		handler.DeleteFornecedor(c, fRepository())
	})

	fornecedorGroup.GET("/:id", func(c *gin.Context) {
		handler.FindFornecedor(c, fRepository())
	})

	fornecedorGroup.PUT("/:id", func(c *gin.Context) {
		handler.UpdateFornecedor(c, fRepository())
	})

	r.GET("/teste", func(c *gin.Context) {
		handler.GetRedis(c, fRepository())
	})
}
