package main

import (
	// "TREINAMENTO-GO/internal/handlers"
	"fmt"
	"net/http"

	// Framework CHI
	"github.com/go-chi/chi/v5"
	"github.com/gustavo-pedro-meira/Treinamento-GO/internal/handlers"
)

func main() {
	// Cria uma nova instancia
	r := chi.NewRouter()

	// Usamos o roteador 'r' para registrar as rotas
	r.Post("/usuario", handlers.HandlerUsuario) // puxo a função do HandlerUsuario do diretorio handlers
	r.Post("/produto", handlers.HandlerProdutos)

	// Nova rota com paramentro URL
	r.Get("/usuario/{nome}", handlers.HandlerBuscarUsuario)
	r.Get("/produto/{nome}", handlers.HandlerBuscarProduto)
	fmt.Println("Servidor rodando na porta 3030 com Chi...")
	http.ListenAndServe(":3030", r) // O segundo parametro agr é o roteador Chi 'r'
}