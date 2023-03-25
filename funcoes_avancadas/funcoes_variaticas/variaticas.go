package main

import "fmt"

func soma(numeros ...int) int {
	// fmt.Println(numeros)
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 123, 43432431, 513342131, 123, 213, 2132, 154, 345)
	fmt.Println(totalDaSoma)
	escrever("Ola Mundo", 10, 20, 3213, 321, 3213, 21321)

}
