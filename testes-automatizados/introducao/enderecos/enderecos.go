package enderecos

import (
	"log"
	"strings"
)

// Tem que ser escrito em maisculo porque ela vai ser exportada
// Verifica se um endereço é valido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	log.Println(enderecoEmLetraMinuscula)
	primeiroPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiroPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiroPalavraDoEndereco)
	}

	return "Tipo Invalido"

}
