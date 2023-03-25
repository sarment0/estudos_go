package main

import "fmt"

func somar(n1 int, n2 int) int {
	return n1 + n2
}

//Aqui podemos trazer mais de un return, até 4 o go permite
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}
	f("texto da funcao 1")

	resultadosCalculos, resultadoSubtracao := calculosMatematicos(10, 15)
	fmt.Println(resultadosCalculos, resultadoSubtracao)
}
