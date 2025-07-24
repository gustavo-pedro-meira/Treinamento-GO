package main

import "fmt"

type Usuario struct {
	ID int
	Nome string
	Email string
	Ativo bool
}

func (u *Usuario) Desativar() {
	u.Ativo = false
	fmt.Printf("O Usuario %s foi Desativado \n", u.Nome)
}

func main() {
	gustavo_user := Usuario {
		ID: 1,
		Nome: "Gustavo",
		Email: "gustavo16pedro@gmail.com",
		Ativo: true,
	}

	fmt.Println("Nome: ", gustavo_user.Nome)
	gustavo_user.Desativar()
	fmt.Println("Esta Ativo? ", gustavo_user.Ativo)
}

