package pagamento

import (
	"cnab/v240/contacorrente"
	"cnab/v240/controle"
	"cnab/v240/empresa"
	"cnab/v240/empresa/endereco"
	"cnab/v240/servico"
	"fmt"
	"strconv"
	"testing"
)

func TestCriarHeaderLote(t *testing.T) {
	// Configuracao do Header do Lote
	pagControle, err := controle.CriarControleHeaderLote(13, 1)
	if err != nil {
		t.Error(err)
	}

	controleErr := controle.Controle{123, 1, controle.HeaderArquivo}

	pagServico, err := servico.CriarServicoHeader(
		servico.Operacao_ARQUIVO_REMESSA,
		servico.Servico_PAGAMENTO_DIVERSOS,
		servico.FormaLancamento_TED_OUTRA_TITULARIDADE)
	if err != nil {
		t.Error(err)
	}

	servicoErr, err := servico.CriarServicoHeader(
		servico.Operacao_ARQUIVO_REMESSA,
		servico.Servico_COBRANCA,
		servico.FormaLancamento_TED_OUTRA_TITULARIDADE)
	if err != nil {
		t.Error(err)
	}

	// Configuracao da empresa
	empInscricao := empresa.Inscricao{1, 123}
	empAg := contacorrente.Agencia{1, "1"}
	empC := contacorrente.Conta{123, "2"}
	empCc := contacorrente.ContaCorrente{empAg, empC, "2"}
	pagEmpresa, err := empresa.CriarEmpresa(empInscricao, "A", empCc, "Minha Empresa")
	if err != nil {
		t.Error(err)
	}

	pagEndEmp, err := endereco.CriarEnderecoEmpresa(
		"Minha Rua",
		123,
		"Casa 0",
		"RIo de Janeiro",
		123,
		"001",
		"RJ",
	)
	if err != nil {
		t.Error(err)
	}

	cenariosTest := []struct {
		controle        controle.Controle
		servico         servico.ServicoHeader
		empresa         empresa.Empresa
		mensagem        string
		enderecoEmpresa endereco.EnderecoEmpresa
		formaPagamento  FormaPagamento
		msgErro         string
	}{
		// Caminho feliz
		{
			pagControle,
			pagServico,
			pagEmpresa,
			"Mensagem da conta",
			pagEndEmp,
			FormaPagamento_DEBITO_CONTACORRENTE,
			"",
		},
		// Controle não condizente com Header de Lote
		{
			controleErr,
			pagServico,
			pagEmpresa,
			"Mensagem da conta",
			pagEndEmp,
			FormaPagamento_DEBITO_CONTACORRENTE,
			"Registro de Controle deve ser equivalente a 1 - HeaderLote",
		},
		// Servico não condizente com Pagamento
		{
			pagControle,
			servicoErr,
			pagEmpresa,
			"Mensagem da conta",
			pagEndEmp,
			FormaPagamento_DEBITO_CONTACORRENTE,
			"Servico deve conter 'PAGAMENTO' no nome da chave",
		},
		// Mensagem com mais de 40 caracteres
		{
			pagControle,
			pagServico,
			pagEmpresa,
			"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX41",
			pagEndEmp,
			FormaPagamento_DEBITO_CONTACORRENTE,
			"Mensagem deve ter até 40 caracteres",
		},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar HeaderLote de pagamento com os parametros do cenario
			_, err := CriarHeaderLotePagamento(
				cenario.controle,
				cenario.servico,
				cenario.empresa,
				cenario.mensagem,
				cenario.enderecoEmpresa,
				cenario.formaPagamento,
			)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.msgErro, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar mensagem\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarHeaderLotePagamento(t *testing.T) {
	// Configuracao do Header do Lote
	pagControle, err := controle.CriarControleHeaderLote(13, 1)
	if err != nil {
		t.Error(err)
	}

	pagServico, err := servico.CriarServicoHeader(
		servico.Operacao_ARQUIVO_REMESSA,
		servico.Servico_PAGAMENTO_DIVERSOS,
		servico.FormaLancamento_TED_OUTRA_TITULARIDADE)
	if err != nil {
		t.Error(err)
	}

	// Configuracao da empresa
	empInscricao := empresa.Inscricao{1, 123}
	empAg := contacorrente.Agencia{1, "1"}
	empC := contacorrente.Conta{123, "2"}
	empCc := contacorrente.ContaCorrente{empAg, empC, "2"}
	pagEmpresa, err := empresa.CriarEmpresa(empInscricao, "A", empCc, "Minha Empresa")
	if err != nil {
		t.Error(err)
	}

	pagEndEmp, err := endereco.CriarEnderecoEmpresa(
		"Minha Rua",
		123,
		"Casa 0",
		"RIo de Janeiro",
		123,
		"001",
		"RJ",
	)
	if err != nil {
		t.Error(err)
	}

	cenariosTest := []struct {
		header   HeaderLotePagamento
		esperado string
	}{
		// Caminho feliz
		{
			HeaderLotePagamento{
				pagControle,
				pagServico,
				" ",
				pagEmpresa,
				"Msg",
				pagEndEmp,
				FormaPagamento_DEBITO_CONTACORRENTE,
				"      ",
				"          "},
			"" + pagControle.Processar() + pagServico.Processar() + " " + pagEmpresa.Processar() + "Msg                                     " + pagEndEmp.Processar() + "01                ",
		},
		// Mensagem com mais de 40 caracteres
		{
			HeaderLotePagamento{
				pagControle,
				pagServico,
				" ",
				pagEmpresa,
				"XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX41",
				pagEndEmp,
				FormaPagamento_DEBITO_CONTACORRENTE,
				"      ",
				"          "},
			"" + pagControle.Processar() + pagServico.Processar() + " " + pagEmpresa.Processar() + "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX" + pagEndEmp.Processar() + "01                ",
		},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar HeaderLotePagamento do cenario
			resultado := cenario.header.Processar()

			// 2. Verificar processamento
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
