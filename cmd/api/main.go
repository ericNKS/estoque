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
	routes.FornecedorRoute(g)
	g.Run()
}
