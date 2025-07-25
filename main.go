package main

import "fmt"

type Vendavel interface {
	ObterPreco() float64
}

type ProdutoFisico struct {
	Nome string
	Preco float64
}

type Assinatura struct {
	Nome string
	PrecoMensal float64
	Meses int
}

func(produto *ProdutoFisico) ObterPreco() float64 {
	return produto.Preco
}

func (assinatura *Assinatura) ObterPreco() float64 {
	return assinatura.PrecoMensal * float64(assinatura.Meses)
}

func CalcularPreco (carrinho []Vendavel) {
	total := 0.0
	for _, item := range carrinho {
		total += item.ObterPreco()
		fmt.Printf("Adicionando item ao carrinho: R$ %.2f\n", item.ObterPreco())
	}
	fmt.Printf("O total do carrinho é R$ %.2f\n", total)
}

func main() {
	notebook := ProdutoFisico {
		Nome: "Notebook",
		Preco: 3500.00,
	}
	curso := Assinatura {
		Nome: "Curso GOLAND",
		PrecoMensal: 12.99,
		Meses: 2,
	}

	meuCarrinho := []Vendavel{&notebook, &curso}
	CalcularPreco(meuCarrinho)
}

