package enderecos

import (
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	// Test + descrição
	enderecoParaTeste := "Avenida Paulista"

	tipoDeEnderecoEsperado := "Avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoEsperado != tipoDeEnderecoRecebido {
		t.Errorf("O tipo recebido é diferente do tipo esperado! Esperava %s e recebeu %s",
			tipoDeEnderecoEsperado,
			tipoDeEnderecoRecebido)
	}
}
