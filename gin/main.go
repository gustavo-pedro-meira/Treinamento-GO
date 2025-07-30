package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Cria umo roteador Gin com config padroes
	r := gin.Default()

	// 2. Define uma rota no metodo GET e qnd algm acessar ela, ele usa a func
	r.GET("/", func(c *gin.Context) {
		// 3. Responde com status 200 e um txt simples
		c.String(200, "Bem - vindo ao GIN")
	})
	r.GET("/bem-vindo", func(c *gin.Context) {
		c.String(200, "Seja mt bem vindo meu caro Gustavo")
	})

	// 4. Inicia o serve na porta padraõ, 8080
	r.Run()
}