package servico

import (
	"fmt"
	"strconv"
	"testing"
)

func TestPodeCriarServico(t *testing.T) {
	servicoHeader, ok := CriarServicoHeader(
		Operacao_LANCAMENTO_CREDITO,
		Servico_PAGAMENTO_SALARIO,
		FormaLancamento_CREDITO_CONTACORRENTE)
	if ok != nil {
		t.Error(ok)
	} else if "C" != servicoHeader.Operacao {
		t.Error("Erro ao configurar a propriedade Operacao do objeto ServicoHeader")
	} else if 30 != servicoHeader.Servico {
		t.Error("Erro ao configurar a propriedade Servico do objeto ServicoHeader")
	} else if 1 != servicoHeader.FormaLancamento {
		t.Error("Erro ao configurar a propriedade FormaLancamento do objeto ServicoHeader")
	}
}

var tabelaErrosCriarServicoHeader = []struct {
	campo           string
	operacao        Operacao
	servico         Servico
	formaLancamento FormaLancamento
	saida           string
}{
	{"Operacao", Operacao("A"), Servico_PAGAMENTO_SALARIO, FormaLancamento_CREDITO_CONTACORRENTE, "Operação não encontrada"},
	{"Servico", Operacao_LANCAMENTO_CREDITO, Servico(100), FormaLancamento_CREDITO_CONTACORRENTE, "Servico não encontrado"},
	{"FormaLancamento", Operacao_LANCAMENTO_CREDITO, Servico_PAGAMENTO_SALARIO, FormaLancamento(100), "Forma de Lançamento não encontrada"},
}

func TestErroAoCriarServicoHeaderComValoresNaoListados(t *testing.T) {
	for _, entry := range tabelaErrosCriarServicoHeader {
		_, ok := CriarServicoHeader(entry.operacao, entry.servico, entry.formaLancamento)
		if ok == nil {
			t.Errorf("Erro, a propriedade %s deve estar listada como constante", entry.campo)
		}
		if entry.saida != fmt.Sprint(ok) {
			t.Errorf("Mensagem de erro não apropriada ao tentar criar ServicoHeader com propriedade %s não listada como constante", entry.campo)
		}
	}
}

func TestPodeProcessarServicoHeader(t *testing.T) {
	servicoHeader := ServicoHeader{
		Operacao_LANCAMENTO_CREDITO,
		Servico_PAGAMENTO_SALARIO,
		FormaLancamento_CREDITO_CONTACORRENTE,
		46,
	}
	resultado := servicoHeader.Processar()
	if len(resultado) != 8 {
		t.Error("Erro no padrão de processamento de ServicoHeader")
	} else if "C3001046" != resultado {
		t.Error("Erro na determinação do resultado ao processar ServicoHeader")
	}
}

var tabelaErrosProcessarServicoHeader = []struct {
	campo   string
	servico ServicoHeader
	saida   string
}{
	{"Operacao", ServicoHeader{Operacao("CC"), Servico_PAGAMENTO_SALARIO, FormaLancamento_CREDITO_CONTACORRENTE, 46}, "C3001046"},
	{"Servico", ServicoHeader{Operacao_LANCAMENTO_CREDITO, Servico(100), FormaLancamento_CREDITO_CONTACORRENTE, 46}, "C1001046"},
	{"FormaLancamento", ServicoHeader{Operacao_LANCAMENTO_CREDITO, Servico_PAGAMENTO_SALARIO, FormaLancamento(100), 46}, "C3010046"},
	{"LayoutLote", ServicoHeader{Operacao_LANCAMENTO_CREDITO, Servico_PAGAMENTO_SALARIO, FormaLancamento_CREDITO_CONTACORRENTE, 9999}, "C3001999"},
}

func TestPodeProcessarServicoHeaderComDigitosAlemDoLimite(t *testing.T) {
	for _, entry := range tabelaErrosProcessarServicoHeader {
		resultado := entry.servico.Processar()
		if len(resultado) != 8 {
			t.Errorf("Erro no padrão de processamento de ServicoHeader quando a propriedade %s tem dígitos além do limite permitido", entry.campo)
		} else if entry.saida != resultado {
			t.Errorf("Erro na determinação do resultado ao processar ServicoHeader quando a propriedade %s tem dígitos além do limite", entry.campo)
		}
	}
}

func TestCriarServicoDetalhe(t *testing.T) {
	cenariosTest := []struct {
		nRegistro  uint32
		seguimento string
		movimento  Movimento
		msgErr     string
	}{
		// Caminho feliz
		{1, "A", Movimento{TipoMovimento_INCLUSAO, CodigoInstrucao_INCLUSAO_REGISTRO_DETALHADO_LIBERADO},
			""},
		// NumeroRegistro lote com mais de 5 dígitos
		{999999, "A", Movimento{TipoMovimento_INCLUSAO, CodigoInstrucao_INCLUSAO_REGISTRO_DETALHADO_LIBERADO},
			"NumeroRegistroLote deve ter até 5 dígitos"},
		// Seguimento com mais de 1 caracter
		{1, "XX", Movimento{TipoMovimento_INCLUSAO, CodigoInstrucao_INCLUSAO_REGISTRO_DETALHADO_LIBERADO},
			"Segmento deve ter até 1 caracter"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar ServicoDetalhe com os parametros do cenario
			_, err := CriarServicoDetalhe(cenario.nRegistro, cenario.seguimento, cenario.movimento)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.msgErr, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar mensagem\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarServicoDetalhe(t *testing.T) {
	cenariosTest := []struct {
		servicoDetalhe ServicoDetalhe
		esperado       string
	}{}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar servicoDetalhe do cenario
			resultado := cenario.servicoDetalhe.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestCriarMovimento(t *testing.T) {
	cenariosTest := []struct {
		tipo   TipoMovimento
		codigo CodigoInstrucao
		msgErr string
	}{
		// Caminho feliz
		{TipoMovimento_INCLUSAO, CodigoInstrucao_INCLUSAO_REGISTRO_DETALHADO_LIBERADO, ""},
		// Tipo com mais de 1 dígito
		{99, CodigoInstrucao_INCLUSAO_REGISTRO_DETALHADO_LIBERADO, "TipoMovimento não encontrado"},
		// Codigo com mais de 2 dígitos
		{TipoMovimento_INCLUSAO, 199, "CodigoInstrucao não encontrado"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar Movimento com os parametros do cenario
			_, err := CriarMovimento(cenario.tipo, cenario.codigo)

			// 2. Verificar mensagem de erro se houver
			if esperado, obtido := cenario.msgErr, fmt.Sprint(err); esperado != "" && esperado != obtido {
				t.Errorf("Erro ao verificar mensagem\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}

func TestProcessarMovimento(t *testing.T) {
	cenariosTest := []struct {
		movimento Movimento
		esperado  string
	}{}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Processar Movimento do cenario
			resultado := cenario.movimento.Processar()

			// 2. Verificar resultado do processamento
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro ao verificar processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
			}
		})
	}
}
