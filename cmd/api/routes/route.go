package routes

import (
	"net/http"

	"github.com/ericNKS/estoque/internal/entities"
	"github.com/ericNKS/estoque/internal/handler"
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/gin-gonic/gin"
)

var v1 = "/api/v1"

func fRepository() *repository.FornecedorRepository {
	fr, err := repository.NewFornecedorRepository()
	if err != nil {
		panic(err)
	}
	return fr
}
func FornecedorRouteV1(r *gin.Engine) {
	fornecedorGroup := r.Group(v1 + "/fornecedor")

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
}

func ProdutoRouteV1(r *gin.Engine) {
	produtoGroup := r.Group(v1 + "/produto")

	produtoGroup.POST("/", func(c *gin.Context) {
		fornecedorId := uint64(12)
		instituicaoId := uint64(1)
		nomeProduto := "Tenis de corrida"
		precoCompra := float32(100)
		precoVenda := float32(299.90)
		qtd := uint16(10)
		p, err := entities.NewProduto(
			fornecedorId,
			instituicaoId,
			nomeProduto,
			precoCompra,
			precoVenda,
			qtd,
		)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"produto": p,
		})

	})
}
