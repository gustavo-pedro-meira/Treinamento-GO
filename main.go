package main

import (
	"github.com/gustavo-pedro-meira/Treinamento-GO/config"
	"github.com/gustavo-pedro-meira/Treinamento-GO/router"
)

func main() {
	// Inicializa as Config
	err := config.Init()
	if err != nil {
		panic(err)
	}

	// Inicializa as Rotas
	router.Initialize()
}