package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Usuario struct {
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func HandlerUsuario(w http.ResponseWriter, r *http.Request) {
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

func HandlerBuscarUsuario(w http.ResponseWriter, r *http.Request) {
	// Chi dar um jeito facil de ler o parametro URL
	nome := chi.URLParam(r, "nome")

	resposta := fmt.Sprintf("Buscando dados do usuario: %s", nome)
	w.Write([]byte(resposta))
}