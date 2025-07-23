package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Printf("Número %d é FizzBuzz \n", i)
		} else if i % 5 == 0 {
			fmt.Printf("Número %d é Buzz \n", i)
		} else if i % 3 == 0 {
			fmt.Printf("Número %d é Fizz \n", i)
		}
	}
}