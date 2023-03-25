package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	var soma2 int16 = numero1 + numero2
	fmt.Println(soma2)

	//atribuicao
	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)

	//operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(!true)
	numero := 3
	numero++
	fmt.Println(numero)
	numero--
	fmt.Println(numero)
	//NAO EXISTE OPERADOR TERNARIO NO GOLANG
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)

}
