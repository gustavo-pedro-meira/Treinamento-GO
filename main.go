package main

import "fmt"

// Interface
type Funcionario interface {
	CalcularSalario() float64
}

// Classes
type Desenvolvedor struct {
	Nome string
	SalarioBase float64
	HorasExtras int
	SalarioFinal float64
}

type Gerente struct {
	Nome string
	SalarioBase float64
	Bonus float64
	SalarioFinal float64
}

// Função
func (d *Desenvolvedor) CalcularSalario() float64 {
	const ValorHoraExtra = 75.50
	d.SalarioFinal = d.SalarioBase + (float64(d.HorasExtras) * ValorHoraExtra)
	return d.SalarioFinal
}

func (g *Gerente) CalcularSalario() float64 {
	g.SalarioFinal = g.SalarioBase + g.Bonus
	return g.SalarioFinal
}

func CalcularCustoFolha(folha []Funcionario) float64 {
	custoTotal := 0.0
	for _, funcionario := range folha {
		custoTotal += float64(funcionario.CalcularSalario())
	}
	return custoTotal
}

func main() {
	dev := Desenvolvedor {
		Nome: "Gustavo",
		SalarioBase: 3000.00,
		HorasExtras: 4,
	}
	ger := Gerente {
		Nome: "Renan",
		SalarioBase: 4000.50,
		Bonus: 1200.12,
	}

	fmt.Println(dev.CalcularSalario())
	fmt.Println(ger.CalcularSalario())

	calculo := []Funcionario{&dev, &ger}

	fmt.Printf("Folha Total: %.2f", CalcularCustoFolha(calculo))
}