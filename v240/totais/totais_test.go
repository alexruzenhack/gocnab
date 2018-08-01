package totais

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCriarTotaisTrailerArquivo(t *testing.T) {
	cenariosTest := []struct {
		qtdLotes, qtdRegistros, qtdContasConcil uint32
		msgErro                                 string
	}{
		// Caminho feliz
		{1, 1, 0, ""},
		// QtdLotes acima de 6 dígitos
		{9999999, 1, 0, "QtdLotes deve ter até 6 dígitos"},
		// QtdRegistros deve ter até 6 dígitos
		{1, 9999999, 0, "QtdRegistros deve ter até 6 dígitos"},
		// QtdContasConcil deve ter até 6 dígitos
		{1, 1, 9999999, "QtdContasConcil deve ter até 6 dígitos"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar Totais com os parametros do cenario
			_, err := CriarTotaisTrailerArquivo(cenario.qtdLotes, cenario.qtdRegistros, cenario.qtdContasConcil)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.msgErro, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Mensagem de erro inapropriada\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarTotaisTrailerArquivo(t *testing.T) {
	cenariosTest := []struct {
		totais    TotaisTrailerArquivo
		resultado string
	}{
		// Caminho feliz
		{TotaisTrailerArquivo{1, 1, 0}, "000001000001000000"},
		// QtdLotes acima de 6 dígitos
		{TotaisTrailerArquivo{9999999, 1, 0}, "999999000001000000"},
		// QtdRegistros acima de 6 dígitos
		{TotaisTrailerArquivo{1, 9999999, 0}, "000001999999000000"},
		// QtdContasConcil acima de 6 dígitos
		{TotaisTrailerArquivo{1, 1, 9999999}, "000001000001999999"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar totais do cenario
			resultado := cenario.totais.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.resultado, resultado; esperado != obtido {
				t.Errorf("Resultado do processamento fora do padrão\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
