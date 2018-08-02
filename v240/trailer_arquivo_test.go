package v240

import (
	"cnab/v240/controle"
	"cnab/v240/totais"
	"fmt"
	"testing"
)

func TestCriarTrailerArquivo(t *testing.T) {
	// Controle
	controle := controle.Controle{1, 1, controle.HeaderArquivo}
	// Totais
	totais := totais.TotaisTrailerArquivo{1, 1, 0}
	_, err := CriarTrailerArquivo(controle, totais)
	if esperado, obtido := "Controle deve possuir o Registro 9 - TrailerArquivo", fmt.Sprint(err); esperado != obtido {
		t.Errorf("Mensagem de erro inadequada\nEsperada: %s\nObtido: %s", esperado, obtido)
	}
}

func TestProcessarTrailerArquivo(t *testing.T) {
	// Controle
	controle := controle.Controle{1, 1, controle.TrailerArquivo}
	// Totais
	totais := totais.TotaisTrailerArquivo{1, 1, 0}
	// TrailerArquivo
	cnabHeader := fmt.Sprintf("%-9s", "")
	cnabTrailer := fmt.Sprintf("%-205s", "")
	trailerArquivo := TrailerArquivo{controle, cnabHeader, totais, cnabTrailer}

	// 1. Processar TrailerArquivo do cenario
	resultado := trailerArquivo.Processar()

	// 2. Verificar resultado do processamento
	if esperado, obtido := 240, len(resultado); esperado != obtido {
		t.Errorf("Erro ao verificar processamento\nEsperado: %d\nObtido: %d", esperado, obtido)
	}
}
