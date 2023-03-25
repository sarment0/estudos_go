package main

import "fmt"

var n int

func init() {
	fmt.Println("Executando a funcao init, tipo um constructor")
	n = 10
}

func main() {
	fmt.Println("funcao main sendo executada")
	fmt.Println(n)
}
