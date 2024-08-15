package main

import (
	"github.com/ericNKS/estoque/internal/repository"
	"github.com/ericNKS/estoque/pkg/handler"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
}
func main() {
	g := gin.Default()
	g.POST("/fornecedor", func(c *gin.Context) {
		fr, err := repository.NewFornecedorRepository()
		if err != nil {
			panic(err)
		}
		handler.CreateFornecedor(c, fr)
	})

	g.GET("/fornecedor", func(c *gin.Context) {
		fr, err := repository.NewFornecedorRepository()
		if err != nil {
			panic(err)
		}
		handler.ListFornecedor(c, fr)
	})

	g.Run()
}
