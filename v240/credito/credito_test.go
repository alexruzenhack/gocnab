package credito

import (
	"fmt"
	"strconv"
	"testing"
)

func TestCriarCredito(t *testing.T) {
	cenariosTest := []struct {
		numeroEmpresa  string
		dataPagamento  uint32
		moeda          Moeda
		valorPagamento uint64
		numeroBanco    string
		errMsg         string
	}{
		// Caminho feliz
		{"123", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 12345, "B123", ""},
		// NumeroEmpresa maior que 20 caracteres
		{"XXXXXXXXXXXXXXXXXXXX21", 1012018, Moeda{TipoMoeda_BRL, 12300000}, 12345, "B123",
			"NumeroEmpresa deve ter até 20 caracteres"},
		// DataPagamento maior que 8 dígitos
		{"123", 101020189, Moeda{TipoMoeda_BRL, 12300000}, 12345, "B123",
			"DataPagamento deve ter até 8 dígitos"},
		// ValorPagamento maior que 15 dígitos
		{"123", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 9999999999999999, "B123",
			"ValorPagamento deve ter até 15 dígitos"},
		// NumeroBanco maior que 20 caracteres
		{"123", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 12345, "BXXXXXXXXXXXXXXXXXXXX",
			"NumeroBanco deve ter até 20 caracteres"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar Credito com os parametros do cenario
			_, err := CriarCredito(
				cenario.numeroEmpresa,
				cenario.dataPagamento,
				cenario.moeda,
				cenario.valorPagamento,
				cenario.numeroBanco,
			)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.errMsg, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar mensagem\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarCredito(t *testing.T) {
	cenariosTest := []struct {
		credito  Credito
		esperado string
	}{
		// Caminho feliz
		{Credito{"123", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 12345, "B123", 0, 0},
			"123                 10102018BRL000000012300000000000000012345B123                00000000000000000000000"},
		// NumeroEmpresa maior que 20 caracteres
		{Credito{"XXXXXXXXXXXXXXXXXXXX21", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 12345, "B123", 0, 0},
			"XXXXXXXXXXXXXXXXXXXX10102018BRL000000012300000000000000012345B123                00000000000000000000000"},
		// DataPagamento maior que 8 dígitos
		{Credito{"123", 101020189, Moeda{TipoMoeda_BRL, 12300000}, 12345, "B123", 0, 0},
			"123                 10102018BRL000000012300000000000000012345B123                00000000000000000000000"},
		// ValorPagamento maior que 15 dígitos
		{Credito{"123", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 9999999999999999, "B123", 0, 0},
			"123                 10102018BRL000000012300000999999999999999B123                00000000000000000000000"},
		// NumeroBanco maior que 20 caracteres
		{Credito{"123", 10102018, Moeda{TipoMoeda_BRL, 12300000}, 12345, "BXXXXXXXXXXXXXXXXXXXX", 0, 0},
			"123                 10102018BRL000000012300000000000000012345BXXXXXXXXXXXXXXXXXXX00000000000000000000000"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar Credito do cenario
			resultado := cenario.credito.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar mensagem do processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestCriarMoeda(t *testing.T) {
	cenariosTest := []struct {
		tipo       TipoMoeda
		quantidade uint64
		errMsg     string
	}{
		// Caminho feliz
		{TipoMoeda_BRL, 12300000, ""},
		// TipoMoeda não encontrado
		{TipoMoeda("XBTC"), 12300000, "TipoMoeda não encontrado"},
		// Quantidade maior que 15 dígitos
		{TipoMoeda_BRL, 9999999999999999, "Quantidade deve ter até 15 dígitos"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar Moeda com os parametros do cenario
			_, err := CriarMoeda(cenario.tipo, cenario.quantidade)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.errMsg, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar mensagem\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarMoeda(t *testing.T) {
	cenariosTest := []struct {
		moeda    Moeda
		esperado string
	}{
		// Caminho feliz
		{Moeda{TipoMoeda_BRL, 12300000}, "BRL000000012300000"},
		// TipoMoeda maior que 3 caracteres
		{Moeda{TipoMoeda("XBTC"), 12300000}, "XBT000000012300000"},
		// Quantidade maior que 15 dígitos
		{Moeda{TipoMoeda_BRL, 9999999999999999}, "BRL999999999999999"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar Moeda do cenário
			resultado := cenario.moeda.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar resultado do processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
