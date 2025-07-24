package main

import "fmt"

func main() {
	type Usuario struct {
		ID int
		Nome string
		Email string
		Ativo bool
	}

	gustavo_user := Usuario {
		ID: 1,
		Nome: "Gustavo",
		Email: "gustavo16pedro@gmail.com",
		Ativo: true,
	}

	fmt.Println("Nome: ", gustavo_user.Nome)
	gustavo_user.Ativo = false
	fmt.Println("Esta Ativo? ", gustavo_user.Ativo)
}