package main

import (
	"fmt"
	"linha-de-comando/app"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	erro := aplicacao.Run(os.Args)
	erro != nil {
		log.Fatal(erro)
	}
}
