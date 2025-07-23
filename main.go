package main

import "fmt"

func main() {
	var soma int = 1

	for soma < 100 {
		soma += soma
		fmt.Printf("Soma atual: %d \n", soma)
	}
}