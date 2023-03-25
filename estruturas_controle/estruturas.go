package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	numero := -10

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que zero.")

	} else if numero < -1 {
		fmt.Println("Numero é menor que -1")
	} else {
		fmt.Println("Numero é menor que zero")
	}
}
