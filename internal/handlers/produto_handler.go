package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Produto struct {
	Nome string `json:"nome"`
	Preco float64 `json:preco`
}

func HandlerProdutos(w http.ResponseWriter, r *http.Request) {
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
// Assinatura padrão de qualquer Handler http, recebendo 2 argumentos
// ResponseWrite usa para retornar dados ao navegador do cliente
// Request é as informações da requisição que o cliente fez (URL, cabeçalhos, corpo da request, etc.)
func HandlerBuscarProduto(w http.ResponseWriter, r *http.Request) {
	// Analisa a requisição 'r' e extrai o valor do paramentro que foi nomeado
	// como 'nome' na definição da rota
	nome := chi.URLParam(r, "nome")

	respota := fmt.Sprintf("Buscando dados do produto: %s", nome)
	// Metodo usado para enviar a resposta final ao navegador
	w.Write([]byte(respota)) // O Write não aceita string, por isso é convertido p byte. Ele esperar um Slices de Bytes
}