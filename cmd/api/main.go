package main

import (
	// "TREINAMENTO-GO/internal/handlers"
	"fmt"
	"log"
	"net/http"

	// Framework CHI
	"github.com/go-chi/chi/v5"
	"github.com/gustavo-pedro-meira/Treinamento-GO/internal/handlers"
	"github.com/gustavo-pedro-meira/Treinamento-GO/internal/database"

)

func main() {
	// Conecta o BD ao iniciar aplicação
	db, err := database.Conectar()
	if err != nil {
		log.Fatalf("Não foi possivel conectar ao BD: %w", err)
	}
	// Garante q a conexão seja fechada ao final da execução main
	defer db.Close()

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")

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