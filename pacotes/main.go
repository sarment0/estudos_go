package main

import (
	"auxiliar/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Alo mundo via pacote")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("123132")
	fmt.Println(erro)
}
