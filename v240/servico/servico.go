package servico

import (
	"errors"
	"fmt"
)

const vLayoutLote uint16 = 46

type ServicoHeader struct {
	// Servico
	// valor padrão igual a LancamentoCredito ('C')
	// *G028
	// tamanho: 1
	Operacao Operacao

	// Quando o serviço adotado é Interoperabilidade entre Contas (23)
	// é obrigatório o preenchimento do campo 18.3C
	// *G025
	// tamanho: 2
	Servico Servico

	// valor padrão igual a 046
	// *G029
	// tamanho: 2
	FormaLancamento FormaLancamento

	// os 2 primeiros digitos sao a versão
	// o último digito representa a release
	// *G030
	// tamanho: 3
	LayoutLote uint16
}

func CriarServicoHeader(operacao Operacao, servico Servico, fLancamento FormaLancamento) (ServicoHeader, error) {
	if _, ok := Operacao_Valor[string(operacao)]; !ok {
		return ServicoHeader{}, errors.New("Operação não encontrada")
	}
	if _, ok := Servico_Valor[uint8(servico)]; !ok {
		return ServicoHeader{}, errors.New("Servico não encontrado")
	}
	if _, ok := FormaLancamento_Valor[uint8(fLancamento)]; !ok {
		return ServicoHeader{}, errors.New("Forma de Lançamento não encontrada")
	}
	return ServicoHeader{
		Operacao:        operacao,
		Servico:         servico,
		FormaLancamento: fLancamento,
		LayoutLote:      vLayoutLote,
	}, nil
}

func (sh ServicoHeader) Processar() string {
	sLayoutLote := fmt.Sprintf("%03d", sh.LayoutLote)
	return sh.Operacao.Processar() + sh.Servico.Processar() + sh.FormaLancamento.Processar() + sLayoutLote[:3]
}

// Operacao trata-se do código alfa usado pela FEBRABAN
// para identificar a transação que será realizada com os
// registros detalhe do lote
// G028
type Operacao string

const (
	Operacao_LANCAMENTO_CREDITO   Operacao = "C"
	Operacao_LANCAMENTO_DEBITO    Operacao = "D"
	Operacao_EXTRATO_CONCILIACAO  Operacao = "E"
	Operacao_EXTRATO_GESTAO_CAIXA Operacao = "G"
	Operacao_INFORMACOES_TITULO   Operacao = "I"
	Operacao_ARQUIVO_REMESSA      Operacao = "R"
	Operacao_ARQUIVO_RETORNO      Operacao = "T"
)

func (o Operacao) Processar() string {
	sOperacao := string(o)
	return sOperacao[:1]
}

var Operacao_Chave = map[string]string{
	"LANCAMENTO_CREDITO":   "C",
	"LANCAMENTO_DEBITO":    "D",
	"EXTRATO_CONCILIACAO":  "E",
	"EXTRATO_GESTAO_CAIXA": "G",
	"INFORMACOES_TITULO":   "I",
	"ARQUIVO_REMESSA":      "R",
	"ARQUIVO_RETORNO":      "T",
}

var Operacao_Valor = map[string]string{
	"C": "LANCAMENTO_CREDITO",
	"D": "LANCAMENTO_DEBITO",
	"E": "EXTRATO_CONCILIACAO",
	"G": "EXTRATO_GESTAO_CAIXA",
	"I": "INFORMACOES_TITULO",
	"R": "ARQUIVO_REMESSA",
	"T": "ARQUIVO_RETORNO",
}

// Servico trata-se do código numérico usado pela FEBRABAN
// para identificar o tipo de serviço / produto (processo)
// contido no arquivo / lote
// G025
type Servico uint8

const (
	Servico_COBRANCA                               Servico = 1
	Servico_BOLETO_PAGAMENTO_ELETRONICO            Servico = 3
	Servico_CONCILIACAO_BANCARIA                   Servico = 4
	Servico_DEBITOS                                Servico = 5
	Servico_CUSTODIA_CHEQUES                       Servico = 6
	Servico_GESTAO_CAIXA                           Servico = 7
	Servico_CONSULTA_MARGEM                        Servico = 8
	Servico_AVERBACAO_CONSIGNACAO_RETENCAO         Servico = 9
	Servico_PAGAMENTO_DIVIDENDOS                   Servico = 10
	Servico_MANUTENCAO_CONSIGNACAO                 Servico = 11
	Servico_CONSIGNACAO_PARCELAS                   Servico = 12
	Servico_GLOSA_CONSIGNACAO                      Servico = 13
	Servico_CONSULTA_TRIBUTOS_APAGAR               Servico = 14
	Servico_PAGAMENTO_FORNECEDOR                   Servico = 20
	Servico_PAGAMENTO_CONTAS_TRIBUTOS_IMPOSTOS     Servico = 22
	Servico_INTEROPERABILIDADE_CONTAS              Servico = 23
	Servico_COMPROR                                Servico = 25
	Servico_COMPROR_ROTATIVO                       Servico = 26
	Servico_ALEGACAO_PAGADOR                       Servico = 29
	Servico_PAGAMENTO_SALARIO                      Servico = 30
	Servico_PAGAMENTO_HONORARIOS                   Servico = 32
	Servico_PAGAMENTO_BOLSA_AUXILIO                Servico = 33
	Servico_PAGAMENTO_PREBENDA                     Servico = 34
	Servico_VENDOR                                 Servico = 40
	Servico_VENDOR_TERMO                           Servico = 41
	Servico_PAGAMENTO_SINISTROS_SEGURO             Servico = 50
	Servico_PAGAMENTO_DESPESAS_VIAJANTE_EMTRANSITO Servico = 60
	Servico_PAGAMENTO_AUTORIZADO                   Servico = 70
	Servico_PAGAMENTO_CREDENCIADOS                 Servico = 75
	Servico_PAGAMENTO_REMUNERACAO                  Servico = 77
	Servico_PAGAMENTO_REPRESENTANTES               Servico = 80
	Servico_PAGAMENTO_BENEFICIOS                   Servico = 90
	Servico_PAGAMENTO_DIVERSOS                     Servico = 98
)

func (s Servico) Processar() string {
	sServico := fmt.Sprintf("%02d", s)
	return sServico[:2]
}

var Servico_Valor = map[uint8]string{
	1:  "COBRANCA",
	3:  "BOLETO_PAGAMENTO_ELETRONICO",
	4:  "CONCILIACAO_BANCARIA",
	5:  "DEBITOS",
	6:  "CUSTODIA_CHEQUES",
	7:  "GESTAO_CAIXA",
	8:  "CONSULTA_MARGEM",
	9:  "AVERBACAO_CONSIGNACAO_RETENCAO",
	10: "PAGAMENTO_DIVIDENDOS",
	11: "MANUTENCAO_CONSIGNACAO",
	12: "CONSIGNACAO_PARCELAS",
	13: "GLOSA_CONSIGNACAO",
	14: "CONSULTA_TRIBUTOS_APAGAR",
	20: "PAGAMENTO_FORNECEDOR",
	22: "PAGAMENTO_CONTAS_TRIBUTOS_IMPOSTOS",
	23: "INTEROPERABILIDADE_CONTAS",
	25: "COMPROR",
	26: "COMPROR_ROTATIVO",
	29: "ALEGACAO_PAGADOR",
	30: "PAGAMENTO_SALARIO",
	32: "PAGAMENTO_HONORARIOS",
	33: "PAGAMENTO_BOLSA_AUXILIO",
	34: "PAGAMENTO_PREBENDA",
	40: "VENDOR",
	41: "VENDOR_TERMO",
	50: "PAGAMENTO_SINISTROS_SEGURO",
	60: "PAGAMENTO_DESPESAS_VIAJANTE_EMTRANSITO",
	70: "PAGAMENTO_AUTORIZADO",
	75: "PAGAMENTO_CREDENCIADOS",
	77: "PAGAMENTO_REMUNERACAO",
	80: "PAGAMENTO_REPRESENTANTES",
	90: "PAGAMENTO_BENEFICIOS",
	98: "PAGAMENTO_DIVERSOS",
}

// FormaLancamento trata-se do código adotado pela FEBRABAN
// para identificar o operação que está contida no lote
type FormaLancamento uint8

const (
	FormaLancamento_CREDITO_CONTACORRENTE                    FormaLancamento = 1
	FormaLancamento_CHEQUE_PAGAMENTO                         FormaLancamento = 2
	FormaLancamento_DOC_TED                                  FormaLancamento = 3
	FormaLancamento_CARTAO_SALARIO                           FormaLancamento = 4 // Somente para servico PatamentoSalarios(30)s(30)
	FormaLancamento_CREDITO_CONTA_POUPANCA                   FormaLancamento = 5
	FormaLancamento_OP_ADISPOSICAO                           FormaLancamento = 10
	FormaLancamento_PAGAMENTO_CONTAS_TRIBUTOS_CODIGODEBARRAS FormaLancamento = 11
	FormaLancamento_TRIBUTO_DARF_NORMAL                      FormaLancamento = 16
	FormaLancamento_TRIBUTO_GPS                              FormaLancamento = 17
	FormaLancamento_TRIBUTO_DARF_SIMPLES                     FormaLancamento = 18
	FormaLancamento_TRIBUTO_IPTU                             FormaLancamento = 19
	FormaLancamento_PAGAMENTO_AUTENTICACAO                   FormaLancamento = 20
	FormaLancamento_TRIBUTO_DARJ                             FormaLancamento = 21
	FormaLancamento_TRIBUTO_GARESP_ICMS                      FormaLancamento = 22
	FormaLancamento_TRIBUTO_GARESP_DR                        FormaLancamento = 23
	FormaLancamento_TRIBUTO_GARESP_ITCMD                     FormaLancamento = 24
	FormaLancamento_TRIBUTO_IPVA                             FormaLancamento = 25
	FormaLancamento_TRIBUTO_LICENCIAMENTO                    FormaLancamento = 26
	FormaLancamento_TRIBUTO_DPVAT                            FormaLancamento = 27
	FormaLancamento_LIQUIDACAO_TITULOS_BANCO                 FormaLancamento = 30
	FormaLancamento_PAGAMENTO_TITULOS_OUTROS_BANCOS          FormaLancamento = 31
	FormaLancamento_EXTRATO_CONTA_CORRENTE                   FormaLancamento = 40
	FormaLancamento_TED_OUTRA_TITULARIDADE                   FormaLancamento = 41
	FormaLancamento_TED_TRANSFERENCIA_CONTA_INVESTIMENTO     FormaLancamento = 44
	FormaLancamento_DEBITO_CONTACORRENTE                     FormaLancamento = 50
	FormaLancamento_EXTRATO_GESTAO_CAIXA                     FormaLancamento = 70
	FormaLancamento_DEPOSITO_JUDICIAL_CONTACORRENTE          FormaLancamento = 71
	FormaLancamento_DEPOSITO_JUDICIAL_POUPANCA               FormaLancamento = 72
	FormaLancamento_EXTRATO_CONTA_INVESTIMENTO               FormaLancamento = 73
)

func (fl FormaLancamento) Processar() string {
	sFormaLancamento := fmt.Sprintf("%02d", fl)
	return sFormaLancamento[:2]
}

var FormaLancamento_Valor = map[uint8]string{
	1:  "CREDITO_CONTACORRENTE",
	2:  "CHEQUE_PAGAMENTO",
	3:  "DOC_TED",
	4:  "CARTAO_SALARIO",
	5:  "CREDITO_CONTA_POUPANCA",
	10: "OP_ADISPOSICAO",
	11: "PAGAMENTO_CONTAS_TRIBUTOS_CODIGODEBARRAS",
	16: "TRIBUTO_DARF_NORMAL",
	17: "TRIBUTO_GPS",
	18: "TRIBUTO_DARF_SIMPLES",
	19: "TRIBUTO_IPTU",
	20: "PAGAMENTO_AUTENTICACAO",
	21: "TRIBUTO_DARJ",
	22: "TRIBUTO_GARESP_ICMS",
	23: "TRIBUTO_GARESP_DR",
	24: "TRIBUTO_GARESP_ITCMD",
	25: "TRIBUTO_IPVA",
	26: "TRIBUTO_LICENCIAMENTO",
	27: "TRIBUTO_DPVAT",
	30: "LIQUIDACAO_TITULOS_BANCO",
	31: "PAGAMENTO_TITULOS_OUTROS_BANCOS",
	40: "EXTRATO_CONTA_CORRENTE",
	41: "TED_OUTRA_TITULARIDADE",
	44: "TED_TRANSFERENCIA_CONTA_INVESTIMENTO",
	50: "DEBITO_CONTACORRENTE",
	70: "EXTRATO_GESTAO_CAIXA",
	71: "DEPOSITO_JUDICIAL_CONTACORRENTE",
	72: "DEPOSITO_JUDICIAL_POUPANCA",
	73: "EXTRATO_CONTA_INVESTIMENTO",
}
