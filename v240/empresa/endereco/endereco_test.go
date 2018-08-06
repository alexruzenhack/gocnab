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
		{"Rua", "Casa", "Rio de Janeiro", "123", "RJ", 1, 21123, ""},
		// Logradouro com mais de 30 caracteres
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "Casa", "Rio de Janeiro", "123", "RJ", 1, 21123, "Logradouro deve ter até 30 caracteres"},
		// Numero com mais de 5 dígitos
		{"Rua", "Casa", "Rio de Janeiro", "123", "RJ", 123456, 21123, "Numero deve ter até 5 dígitos"},
		// Complemento com mais de 15 caracteres
		{"Rua", "XXXXXXXXXXXXXXXX", "Rio de Janeiro", "123", "RJ", 1, 21123, "Complemento deve ter até 15 caracteres"},
		// Cidade com mais de 20 caracteres
		{"Rua", "Casa", "XXXXXXXXXXXXXXXXXXXXX", "123", "RJ", 1, 21123, "Cidade deve ter até 20 caracteres"},
		// Cep com mais de 5 dígitos
		{"Rua", "Casa", "Rio de Janeiro", "123", "RJ", 1, 123456, "Cep deve ter até 5 dígitos"},
		// ComplementoCep com mais de 3 caracteres
		{"Rua", "Casa", "Rio de Janeiro", "1234", "RJ", 1, 21123, "ComplementoCep deve ter até 3 caracteres"},
		// Estado com mais de 2 caracteres
		{"Rua", "Casa", "Rio de Janeiro", "123", "RJZ", 1, 21123, "Estado deve ter até 2 caracteres"},
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

func TestProcessarEnderecoEmpresa(t *testing.T) {
	cenariosTest := []struct {
		endereco EnderecoEmpresa
		esperado string
	}{
		// Caminho feliz
		{EnderecoEmpresa{"RUA", 123, "CASA", "RIO", 21123, "001", "RJ"},
			"RUA                           00123CASA           RIO                 21123001RJ"},
		// Logradouro com mais de 30 caracteres
		{EnderecoEmpresa{"RUAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", 123, "CASA", "RIO", 21123, "001", "RJ"},
			"RUAXXXXXXXXXXXXXXXXXXXXXXXXXXX00123CASA           RIO                 21123001RJ"},
		// Numero com mais de 5 dígitos
		{EnderecoEmpresa{"RUA", 123456, "CASA", "RIO", 21123, "001", "RJ"},
			"RUA                           12345CASA           RIO                 21123001RJ"},
		// Complemento com mais de 15 caracteres
		{EnderecoEmpresa{"RUA", 123, "CASAXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "RIO", 21123, "001", "RJ"},
			"RUA                           00123CASAXXXXXXXXXXXRIO                 21123001RJ"},
		// Cidade com mais de 20 caracteres
		{EnderecoEmpresa{"RUA", 123, "CASA", "RIOXXXXXXXXXXXXXXXXXXXX", 21123, "001", "RJ"},
			"RUA                           00123CASA           RIOXXXXXXXXXXXXXXXXX21123001RJ"},
		// Cep com mais de 5 dígitos
		{EnderecoEmpresa{"RUA", 123, "CASA", "RIO", 211234, "001", "RJ"},
			"RUA                           00123CASA           RIO                 21123001RJ"},
		// ComplementoCep com mais de 3 caracteres
		{EnderecoEmpresa{"RUA", 123, "CASA", "RIO", 21123, "1234", "RJ"},
			"RUA                           00123CASA           RIO                 21123123RJ"},
		// Estado com mais de 2 caracteres
		{EnderecoEmpresa{"RUA", 123, "CASA", "RIO", 21123, "001", "RJZ"},
			"RUA                           00123CASA           RIO                 21123001RJ"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar endereco fornecido no cenario
			resultado := cenario.endereco.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
