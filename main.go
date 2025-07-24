package main

import "fmt"

type ContaBancaria struct {
	Titular string
	Saldo float64
}

func (conta *ContaBancaria) Depositar(valor float64) {
	if valor < 0.0 {
		fmt.Println("Valor negativo \n")
	} else {
		conta.Saldo += valor
		fmt.Printf("Valor depositado: %.2f \n", valor)
	}
}

func (conta *ContaBancaria) Sacar(valor float64) {
	if valor > conta.Saldo {
		fmt.Println("Saldo Insuficiente")
	} else {
		conta.Saldo -= valor
		fmt.Printf("Valor Sacado: %.2f \n", valor)
	}
}

func (conta *ContaBancaria) VerSaldo() {
	fmt.Printf("Saldo Disponivel: %.2f \n", conta.Saldo)
}

func main() {
	conta := ContaBancaria {
		Titular: "Gustavo",
		Saldo: 100.00,
	}

	conta.Depositar(250)
	conta.VerSaldo()
	conta.Sacar(400.99)
	conta.Sacar(300.27)
	conta.VerSaldo()
}

