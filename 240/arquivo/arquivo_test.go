package arquivo

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCriarArquivo(t *testing.T) {
	cenariosTest := []struct {
		sequencia uint32
		densidade Densidade
		msgErro   string
	}{
		{1, Densidade_1600_BPI, ""},
		{0, Densidade_1600_BPI, "Sequencia deve ser maior do que zero"},
		{9999999, Densidade_1600_BPI, "Sequencia deve ter até 6 dígitos"},
		{1, Densidade(1), "Valor de Densidade não permitido"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar Arquivo com os parametros do cenário
			_, err := CriarArquivo(cenario.sequencia, cenario.densidade)

			// 2. Verificar mensagem de erro caso tenha
			if esperado, obtido := cenario.msgErro, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao criar Arquivo\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarCodigo(t *testing.T) {
	cenariosTest := []struct {
		codigo    Codigo
		resultado string
	}{
		{Codigo_REMESSA, "1"},
		{Codigo(10), "1"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar resultado do cenario
			resultado := cenario.codigo.processar()

			// 2. Verificar tamanho do resultado
			if esperado, obtido := 1, len(resultado); esperado != obtido {
				t.Errorf("Erro ao verificar tamanho do Codigo processado\nEsperado: %d\nObtido: %d", esperado, obtido)
			}

			// 3. Verificar padrão do resultado
			if esperado, obtido := cenario.resultado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar padrão do Codigo processado\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarDensidade(t *testing.T) {
	cenariosTest := []struct {
		densidade Densidade
		resultado string
	}{
		{Densidade_1600_BPI, "01600"},
		{Densidade(123456), "12345"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar resultado do cenario
			resultado := cenario.densidade.processar()

			// 2. Verificar tamanho do resultado
			if esperado, obtido := 5, len(resultado); esperado != obtido {
				t.Errorf("Erro ao verificar tamanho do tipo Densidade processado\nEsperado: %d\nObtido: %d", esperado, obtido)
			}

			// 3. Verificar padrão do resultado
			if esperado, obtido := cenario.resultado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar padrão do tipo Densidade processado\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarArquivo(t *testing.T) {
	var vLayoutArquivo uint16 = 103

	cenariosTest := []struct {
		arquivo   Arquivo
		resultado string
	}{
		{Arquivo{Codigo_REMESSA, 11121991, 162000, 1, vLayoutArquivo, Densidade_1600_BPI}, "11112199116200000000110301600"},
		{Arquivo{Codigo_REMESSA, 111219910, 162000, 1, vLayoutArquivo, Densidade_1600_BPI}, "11112199116200000000110301600"},
		{Arquivo{Codigo_REMESSA, 11121991, 16200099, 1, vLayoutArquivo, Densidade_1600_BPI}, "11112199116200000000110301600"},
		{Arquivo{Codigo_REMESSA, 11121991, 162000, 1234567, vLayoutArquivo, Densidade_1600_BPI}, "11112199116200012345610301600"},
		{Arquivo{Codigo_REMESSA, 11121991, 1620009, 1, 1234, Densidade_1600_BPI}, "11112199116200000000112301600"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar resultado do cenario
			resultado := cenario.arquivo.Processar()

			// 2. Verificar tamanho do resultado
			if esperado, obtido := 29, len(resultado); esperado != obtido {
				t.Errorf("Erro ao verificar tamanho do tipo Arquivo processado\nEsperado: %d\nObtido: %d", esperado, obtido)
			}

			// 3. Verificar padrão do resultado
			if esperado, obtido := cenario.resultado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar padrão do tipo Arquivo processado\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
