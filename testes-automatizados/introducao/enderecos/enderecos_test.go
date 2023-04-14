// teste de unidade
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		{"Praça das Rosas", "Tipo Invalido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		{"", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
	//
	// enderecoParaTeste := "Avenida Paulista"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo reebido eh diferente do esperado! Esperava %s e recebeu %s",
	// 		tipoDeEnderecoEsperado,
	// 		tipoDeEnderecoRecebido,
	// 	)
	// }
}

//go test --cover
//go test ./... --cover
func TestQualquer(t *testing.T) {
	t.Parallel() //para rodar testes em paralelo
	if 1 > 2 {
		t.Errorf("Teste quebrou.")
	}
}
