package main

import (
	"errors"
	"fmt"
)

func main() {
	numero := -1000
	fmt.Println(numero)

	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	numero4 := 12322222222.40
	fmt.Println(numero4)

	var erro error = errors.New("Erro interno na bagaça.")
	fmt.Println(erro)
}
