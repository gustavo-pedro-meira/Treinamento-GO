package main

import "fmt"

func main() {
	var comidas []string

	comidas = append(comidas, "Arroz")
	comidas = append(comidas, "Carne")
	comidas = append(comidas, "Feijão")

	fmt.Println(comidas)

	prato := comidas[1:2]
	fmt.Println(prato)
}