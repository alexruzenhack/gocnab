package pagamento

import (
	"cnab/v240/controle"
	"cnab/v240/empresa"
	"cnab/v240/empresa/endereco"
	"cnab/v240/servico"
	"errors"
	"fmt"
	"strings"
)

type HeaderLotePagamento struct {
	// G001, *G002, *G003
	// tamanho: 8
	Controle controle.Controle

	// *G028, *G025, *G029, *G030
	// tamanho: 8
	Servico servico.ServicoHeader

	// Texto de observações destinado para uso exclusivo da FEBRABAN
	// preencher com brancos
	// G004
	// tamanho: 1
	CnabHeader string

	// *G005, *G006, *G007, *G008, *G009, *G010, *G011, *G012, G013
	// tamanho: 85
	Empresa empresa.Empresa

	// Texto referente a mensagens que serão impressas nos documentos
	// e/ou avisos a serem emitidos
	// Informação 1: Mensagem genérica
	// *G031
	// tamanho: 40
	Mensagem string

	// G032, G033, G034, G035, G036
	// tamanho: 80
	EnderecoEmpresa endereco.EnderecoEmpresa

	// Possibilitar ao Pagador, mediante acordo com seu Banco de Relacionamento
	// a forma de pagamento do compromisso
	// P014
	// tamanho: 2
	FormaPagamento FormaPagamento

	// G004
	// tamanho: 6
	CnabTrailer string

	// Código adotado pela FEBRABAN para identificar as ocorrências
	// detectadas no processamento
	// Pode-se informar até 5 ocorrências simultaneamente, cada uma delas
	// codificada com dois dígitos
	// *G059
	// tamanho: 10
	Ocorrencias string
}

type FormaPagamento int

const (
	FormaPagamento_DEBITO_CONTACORRENTE            FormaPagamento = 1
	FormaPagamento_DEBITO_EMPRESTIMO_FINANCIAMENTO FormaPagamento = 2
	FormaPagamento_DEBITO_CARTAO_CREDITO           FormaPagamento = 3
)

func CriarHeaderLotePagamento(ctrl controle.Controle, svco servico.ServicoHeader, empsa empresa.Empresa, msg string, endcoEmpsa endereco.EnderecoEmpresa, frmaPag FormaPagamento) (HeaderLotePagamento, error) {
	if ctrl.Registro != 1 {
		return HeaderLotePagamento{}, errors.New("Registro de Controle deve ser equivalente a 1 - HeaderLote")
	} else if !strings.HasPrefix(servico.Servico_Valor[uint8(svco.Servico)], "PAGAMENTO") {
		return HeaderLotePagamento{}, errors.New("Servico deve conter 'PAGAMENTO' no nome da chave")
	} else if len(msg) > 40 {
		return HeaderLotePagamento{}, errors.New("Mensagem deve ter até 40 caracteres")
	}

	CnabHeader := " "
	CnabTrailer := fmt.Sprintf("%-6s", "")
	Ocorrencias := fmt.Sprintf("%-10s", "")

	return HeaderLotePagamento{
		ctrl,
		svco,
		CnabHeader,
		empsa,
		msg,
		endcoEmpsa,
		frmaPag,
		CnabTrailer,
		Ocorrencias,
	}, nil
}

func (hlp HeaderLotePagamento) Processar() string {
	sMensagem := fmt.Sprintf("%-40s", hlp.Mensagem)
	sFormaPagamento := fmt.Sprintf("%02d", hlp.FormaPagamento)
	sCnabTrailer := fmt.Sprintf("%-6s", "")
	sOcorrencias := fmt.Sprintf("%-10s", "")
	return hlp.Controle.Processar() + hlp.Servico.Processar() + hlp.CnabHeader[:1] + hlp.Empresa.Processar() + sMensagem[:40] + hlp.EnderecoEmpresa.Processar() + sFormaPagamento[:2] + sCnabTrailer[:6] + sOcorrencias[:10]
}
