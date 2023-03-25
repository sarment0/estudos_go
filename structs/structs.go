package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")

	var u usuario
	u.idade = 36
	u.nome = "Fabio"
	fmt.Println(u)

	//setters direto no construtor
	enderecoExemplo := endereco{"Rua Pereira Barreto", 159}
	usuario2 := usuario{"Fabio", 36, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Fabio"}
	fmt.Println(usuario3)
}
