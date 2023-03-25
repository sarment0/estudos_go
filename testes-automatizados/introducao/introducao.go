package main

import (
	"fmt"
	"testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("avenida Paulista")
	fmt.Println(tipoEndereco)
}
