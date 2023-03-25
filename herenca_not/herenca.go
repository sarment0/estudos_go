package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herença")
	p1 := pessoa{"Fabio", "Sarmento", 36, 175}
	fmt.Println(p1)
	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome)
}
