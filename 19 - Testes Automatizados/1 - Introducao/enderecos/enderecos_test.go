package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTest := "Vila do conde"

	tipoDeEnderecoEsperado := "Vila"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTest)
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido e diferente do esperado! esperava %s e recebeu %s", tipoDeEnderecoEsperado, tipoDeEnderecoRecebido)
	}
}
