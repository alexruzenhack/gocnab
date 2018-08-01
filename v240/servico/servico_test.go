package servico

import (
	"fmt"
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
