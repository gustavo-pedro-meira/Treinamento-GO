package main

import "fmt"

func main() {
	nome_empresa := "Ambev"
	const ano_fundacao = 2004
	var num_funcionarios int = 24000
	var media_salarial float64 = 1800.50
	trabalho_remoto := false

	if trabalho_remoto == true {
		fmt.Printf(" O trabalho na %s é remoto! \n", nome_empresa)
	} else {
		fmt.Printf(" O trabalho na %s não é remoto! \n", nome_empresa)
	}

	fmt.Printf(" Empresa: %s \n Fundado em: %d \n Número de Funcionários: %d \n Média Salarial: %.2f \n Oferece remoto? %t", nome_empresa, ano_fundacao, num_funcionarios, media_salarial, trabalho_remoto)
}