package main

import "fmt"

func main() {
	if nome, idade := "Gustavo", 18; nome == "Gustavo"{
		fmt.Printf("Nome correto, idade: %d", idade)
	} else {
		fmt.Println("Nome Incorreto")
	}
}