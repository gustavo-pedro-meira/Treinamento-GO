package main

import "fmt"

func main() {
	for i := 0; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Printf("Número %d é par \n", i)
		} 
	}
}