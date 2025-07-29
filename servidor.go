package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerRaiz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ola mundo! Meu primeiro server GOLANG está no ar.")
}

func handlerOla(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		nome = "Visitante"
	}
	fmt.Fprintf(w, "Opa, %s! Bem vindo a minha API", nome)
}

func handlerSobre(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	if nome == "" {
		nome = "Visitante"
	}
	fmt.Fprintf(w, "Este é um servidor web feito por %s", nome)
}

func main() {
	http.HandleFunc("/", handlerRaiz)
	http.HandleFunc("/ola", handlerOla)
	http.HandleFunc("/sobre", handlerSobre)

	fmt.Println("Servidor escutando na porta 8080...")
	fmt.Println("Acesse http://localhost:3030 ou http://localhost:3030/ola")

	log.Fatal(http.ListenAndServe(":3030", nil))
}