package main

import (
	"github.com/ericNKS/estoque/cmd/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load("../../.env")
}
func main() {
	g := gin.Default()
	routes.FornecedorRouteV1(g)
	routes.ProdutoRouteV1(g)
	g.Run("127.0.0.1:8080")
}
