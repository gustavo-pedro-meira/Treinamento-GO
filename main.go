package main

import "fmt"

type Produto struct {
	Nome string
	Preco float64
	Quantidade int
}

func (p *Produto) AplicarDesconto(percentual float64) {
	desconto := p.Preco * (percentual / 100)
	valor_desconto := p.Preco - desconto
	fmt.Printf("Valor do Produto com Desconto: %.2f \n", valor_desconto)
}

func (p *Produto) ValorTotalEstoque() {
	valor_total := p.Preco * float64(p.Quantidade)
	fmt.Println("Valor total em estoque: ", valor_total)
}

func (p *Produto) AdicionarEstoque(quantidade int) {
	p.Quantidade += quantidade
	fmt.Printf("Quantidade de %d aumentada para o produto %s \n", p.Quantidade, p.Nome)
}

func(p *Produto) VenderProduto(quantidade int) {
	if quantidade > p.Quantidade {
		fmt.Println("Estoque Insuficiente")
	} else {
		p.Quantidade -= quantidade
		fmt.Println("Venda Realizada")
	}
}

func main() {
	notebook := Produto {
		Nome: "Samsung Book",
		Preco: 5000.00,
		Quantidade: 5,
	}

	notebook.ValorTotalEstoque()
	notebook.AplicarDesconto(10)
	notebook.AdicionarEstoque(2)
	notebook.VenderProduto(2)
}

