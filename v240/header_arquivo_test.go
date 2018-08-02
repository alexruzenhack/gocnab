package v240

import (
	"cnab/v240/arquivo"
	"cnab/v240/contacorrente"
	"cnab/v240/controle"
	"cnab/v240/empresa"
	"fmt"
	"strconv"
	"testing"
)

func TestCriarHeaderArquivo(t *testing.T) {
	cenariosTest := []struct {
		nomeBanco, reservadoEmpresa, msgErro string
	}{
		// Caminho feliz
		{"Meu Banco", "Nota da Empresa", ""},
		// NomeBanco acima de 30 dígitos
		{"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX", "Nota da Empresa", "NomeBanco deve ter até 30 dígitos"},
		// ReservadoEmpresa acima de 20 dígitos
		{"Meu Banco", "XXXXXXXXXXXXXXXXXXXXX", "ReservadoEmpresa deve ter até 20 dígitos"},
	}

	ctrle := controle.Controle{}
	empsa := empresa.Empresa{}
	arqvo := arquivo.Arquivo{}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar HeaderArquivo com os parametros do cenário
			_, err := CriarHeaderArquivo(ctrle, empsa, arqvo, cenario.nomeBanco, cenario.reservadoEmpresa)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.msgErro, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar o cenário\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarHeaderArquivo(t *testing.T) {
	// Configuração padrão de HeaderArquivo
	cnabHeader := fmt.Sprintf("%-9s", "")
	cnabBody := fmt.Sprintf("%-10s", "")
	cnabTrailer := fmt.Sprintf("%-29s", "")
	reservadoBanco := fmt.Sprintf("%-20s", "")

	// Configuração padrão de Controle em HeaderArquivo
	ctrle, _ := controle.CriarControleHeaderArquivo(uint16(1))

	// Configuração padrão de Empresa
	inscricao := empresa.Inscricao{1, 123}
	ag := contacorrente.Agencia{1, "1"}
	ca := contacorrente.Conta{123, "2"}
	cc := contacorrente.ContaCorrente{ag, ca, "2"}
	empsa := empresa.Empresa{inscricao, "A", cc, "Nome Minha Empresa"}

	// Configuração padrão de Arquivo
	arqvo := arquivo.Arquivo{arquivo.Codigo_REMESSA, 11121991, 162000, 1, uint16(103), arquivo.Densidade_1600_BPI}

	cenariosTest := []struct {
		headerArquivo HeaderArquivo
		resultado     string
	}{
		// Caminho feliz
		{HeaderArquivo{
			ctrle,
			cnabHeader,
			empsa,
			"nomebanco",
			cnabBody,
			arqvo,
			reservadoBanco,
			"reservadoempresa",
			cnabTrailer},
			ctrle.Processar() + cnabHeader + empsa.Processar() + "nomebanco                     " + cnabBody + arqvo.Processar() + reservadoBanco + "reservadoempresa    " + cnabTrailer},
		// NomeBanco acima de 30 dígitos
		{HeaderArquivo{
			ctrle,
			cnabHeader,
			empsa,
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			cnabBody,
			arqvo,
			reservadoBanco,
			"reservadoempresa",
			cnabTrailer},
			ctrle.Processar() + cnabHeader + empsa.Processar() + "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" + cnabBody + arqvo.Processar() + reservadoBanco + "reservadoempresa    " + cnabTrailer},
		// ReservadoEmpresa acima de 20 dígitos
		{HeaderArquivo{
			ctrle,
			cnabHeader,
			empsa,
			"nomebanco",
			cnabBody,
			arqvo,
			reservadoBanco,
			"XXXXXXXXXXXXXXXXXXXXX",
			cnabTrailer},
			ctrle.Processar() + cnabHeader + empsa.Processar() + "nomebanco                     " + cnabBody + arqvo.Processar() + reservadoBanco + "XXXXXXXXXXXXXXXXXXXX" + cnabTrailer},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar headerArquivo do cenario
			resultado := cenario.headerArquivo.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.resultado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar o resultado do cenário\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
