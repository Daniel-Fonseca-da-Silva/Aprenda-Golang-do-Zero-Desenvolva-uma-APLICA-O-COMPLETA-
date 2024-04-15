package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua do Marques", "Rua"},
		{"Avenida do carmo", "Avenida"},
		{"Parque geraldo", "Tipo invalido"},
		{"Estrada", "Rua"},
		{"RUA DOS BOBOS", "Rua"},
		{enderecoInserido: "", retornoEsperado: "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s e diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}

}

func TestQualquerCoisa(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}

}
