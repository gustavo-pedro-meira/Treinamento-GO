package main

import "fmt"

func main() {
	numerospares := [3]int{2, 4, 6}
	fmt.Println(numerospares)
	numerospares[2] = 8
	fmt.Println(numerospares[2])
}