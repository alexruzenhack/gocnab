package endereco

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCriarEnderecoEmpresa(t *testing.T) {
	cenariosTest := []struct {
		Logradouro, Complemento, Cidade, ComplementoCep, Estado string
		Numero, Cep                                             uint32
		msgErro                                                 string
	}{
		// Caminho feliz
		{},
		// Logradouro com mais de 30 caracteres
		{},
		// Numero com mais de 5 dígitos
		{},
		// Complemento com mais de 15 caracteres
		{},
		// Cidade com mais de 20 caracteres
		{},
		// Cep com mais de 5 dígitos
		{},
		// ComplementoCep com mais de 3 caracteres
		{},
		// Estado com mais de 2 caracteres
		{},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar EnderecoEmpresa com os parametros do cenario
			_, err := CriarEnderecoEmpresa(
				cenario.Logradouro,
				cenario.Numero,
				cenario.Complemento,
				cenario.Cidade,
				cenario.Cep,
				cenario.ComplementoCep,
				cenario.Estado,
			)

			// 2. Verificar mensagem de erro
			if esperado, obtido := cenario.msgErro, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar mensagem\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
