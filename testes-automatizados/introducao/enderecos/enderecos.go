package enderecos

import "strings"

//Tem que ser escrito em maisculo porque ela vai ser exportada
//Verifica se um endereço é valido
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Estrada", "Rodovia"}
	enderecoEmLetraMinuscula := strings.ToLower(endereco)
	primeiroPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]
	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiroPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiroPalavraDoEndereco
	}

	return "Tipo Invalido"

}
