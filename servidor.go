package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Produto struct {
	Nome string `json:"nome"`
	Preco float64 `json:preco`
}

type Usuario struct {
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func handlerUsuario(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo HTTP não Permitido", http.StatusMethodNotAllowed)
		return
	}

	var usuario Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, "Erro ao Decodificar o JSON", http.StatusMethodNotAllowed)
		return
	}

	fmt.Printf("Usuario '%+v' Recebido", usuario)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Usuario %s cadastrado com sucesso!", usuario.Nome)
}

func handlerProdutos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Metodo HTTP não Permitido", http.StatusMethodNotAllowed)
		return
	}

	var produto Produto

	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, "Erro ao decodificar o JSON", http.StatusMethodNotAllowed)
		return
	}

	fmt.Printf("Produto recebido: %+v\n", produto)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Produto '%s' recebido com sucesso!", produto.Nome)
}

func main() {
	http.HandleFunc("/produtos", handlerProdutos)
	http.HandleFunc("/usuario", handlerUsuario)
	fmt.Println("Servidor rodando na porta 3030...")
	log.Fatal(http.ListenAndServe(":3030", nil))

}