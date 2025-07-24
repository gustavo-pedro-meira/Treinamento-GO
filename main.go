package main

import "fmt"

func main() {
	notaAlunos := make(map[string]float64)

	notaAlunos["Gustavo"] = 9.9
	notaAlunos["Renan"] = 8.7

	fmt.Printf("A nota de Gustavo é %.1f \n", notaAlunos["Gustavo"])
	fmt.Println(notaAlunos)

	nota, existe := notaAlunos["Gustavo"]
	if existe {
		fmt.Println("A nota Existe", nota)
	} else {
		fmt.Println("A nota não Existe")
	}
}