package main

import "fmt"

func closure() func() {
	texto := "Dentro da funcao closure"
	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da funacao main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
