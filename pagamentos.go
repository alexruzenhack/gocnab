package cnab

import (
	"cnab/controle"
)

// Header
type Header struct {
	Controle controle.Controle

	// Servico
	// valor padrão igual a LancamentoCredito ('C')
	// tamanho: 1
	Operacao Operacao // *G028

	// Quando o serviço adotado é Interoperabilidade entre Contas (23)
	// é obrigatório o preenchimento do campo 18.3C
	// tamanho: 2
	Servico Servico // *G025

	// valor padrão igual a 046
	// tamanho: 2
	FormaLancamento Lancamento // *G029

	// os 2 primeiros digitos sao a versão
	// o último digito representa a release
	// tamanho: 3
	LayoutLote int // *G030

	// CNAB
	// valor padrão caracter vazio ' '
	// tamanho: 1
	Cnab string // G004

	// Empresa
	// Preenchimento deste campo é obrigatório para DOC e TED
	// formas de Lançamento: 03, 41, 43
	// tamanho: 1
	TipoInscricao int // *G005

	// Número de inscrição da Empresa ou Pessoa Física
	// perante instituição governamental
	// Preencher com zeros caso não seja informado
	// tamanho: 14
	NumeroInscricao int // *G006

	// Código adotado pelo Banco responsável pela conta
	// para identificar o Contrato entre este e a Empersa Cliente
	// tamanho: 20
	Convenio string // *G007

	// Conta Corrente
	// Código adotado pelo Banco responsável pela conta
	// para identificar qual unidade está vinculada a conta corrente
	// tamanho: 5
	CodigoAgencia int // *G008

	// Conta adotado pelo Banco responsável pela conta corrente
	// para verificação da autenticiade do código da Agência
	// tamanho: 1
	DvAgencia string // *G009

	// Número adotado pelo Banco, para identificar univocamente
	// a conta corrente utilizada pelo cliente
	// tamanho: 12
	NumeroConta int // *G010

	// Código adotado pelo Banco responsável pela conta corrente
	// para verificação da autenticidade do Número da Conta Corrente
	// tamanho: 1
	DvConta string // *G011

	// Código adotado pelo Banco responsável pela conta corrente
	// para verificação da autenticidade do par Código da Agência
	// Número da Conta Corrente
	// tamanho: 1
	DvContaCorrente string // *G012

	// Nome que identifica a pessoa, física ou jurídica,
	// a qual se quer fazer referencia
	// tamanho: 30
	Nome string // G012

	// Texto referente a mensagens que serão impressas nos documentos
	// e/ou avisos a serem emitidos
	// Informação 1: Mensagem genérica
	// tamanho: 40
	Mensagem string // *G031

	// Endereço da Empresa
	// Texto referente a localização da rua/avenida, número, complemento,
	// e bairro utilizado para entrega de correspondencia. Utilizado
	// também para endereço de e-mail para entrega eletrônica de informação
	// e para número de celular para envio de mensagem SMS
	// tamanho: 30
	Logradouro string // G032

	// tamanho: 5
	Numero int // G032

	// tamanho: 15
	Complemento string // G032

	// Texto referente ao nome do município componente do endereço
	// utilizado para entrega de correspondencia
	// tamanho: 20
	Cidade string // G033

	// Código adotado pela EBCT (Empresa Brasileira de Correios e Telégrafos)
	// para identificação de logradouros
	// tamanho: 5
	Cep int // G034

	// Código para complementação do código do CEP
	// tamanho: 3
	ComplementoCep string // G035

	// Código do estado, unidade da federação componente do endereço
	// utilizado para entrega de correspondência
	// tamanho: 2
	Estado string // G036

	// Possibilitar ao Pagador, mediante acordo com seu Banco de Relacionamento
	// a forma de pagamento do compromisso
	// tamanho: 2
	FormaPagamento Pagamento // P014
	// Cnab
	// Cnab        int // G004

	// Código adotado pela FEBRABAN para identificar as ocorrências
	// detectadas no processamento
	// Pode-se informar até 5 ocorrências simultaneamente, cada uma delas
	// codificada com dois dígitos
	// tamanho: 10
	Ocorrencias int // *G059
}

// Operacao trata-se do código alfa usado pela FEBRABAN
// para identificar a transação que será realizada com os
// registros detalhe do lote
// G028
type Operacao string

const (
	LancamentoCredito  Operacao = "C"
	LancamentoDebito   Operacao = "D"
	ExtratoConciliacao Operacao = "E"
	ExtratoGestaoCaixa Operacao = "G"
	InformacoesTitulo  Operacao = "I"
	ArquivoRemessa     Operacao = "R"
	ArquivoRetorno     Operacao = "T"
)

// Servico trata-se do código numérico usado pela FEBRABAN
// para identificar o tipo de serviço / produto (processo)
// contido no arquivo / lote
// G025
type Servico int

const (
	Cobranca                            Servico = 1
	BoletoPagamentoEletronico           Servico = 3
	ConciliacaoBancaria                 Servico = 4
	Debitos                             Servico = 5
	CustodiaCheques                     Servico = 6
	GestaoCaixa                         Servico = 7
	ConsultaMargem                      Servico = 8
	AverbacaoConsignacaoRetencao        Servico = 9
	PagamentoDividendos                 Servico = 10
	ManutencaoConsignacao               Servico = 11
	ConsignacaoParcelas                 Servico = 12
	GlosaConsignacao                    Servico = 13
	ConsultaTributosAPagar              Servico = 14
	PagamentoFornecedor                 Servico = 20
	PagamentoContasTributosImpostos     Servico = 22
	InteroperabilidadeContas            Servico = 23
	Compror                             Servico = 25
	ComprorRotativo                     Servico = 26
	AlegacaoPagador                     Servico = 29
	PagamentoSalarios                   Servico = 30
	PagamentoHonorarios                 Servico = 32
	PagamentoBolsaAuxilio               Servico = 33
	PagamentoPrebenda                   Servico = 34
	Vendor                              Servico = 40
	VendorTermo                         Servico = 41
	PagamentoSinistrosSegurados         Servico = 50
	PagamentoDespesasViajanteEmTransito Servico = 60
	PagamentoAutorizado                 Servico = 70
	PagamentoCredenciados               Servico = 75
	PagamentoRemuneracao                Servico = 77
	PagamentoRepresentantes             Servico = 80
	PagamentoBeneficios                 Servico = 90
	PagamentosDiversos                  Servico = 98
)

// Lancamento trata-se do código adotado pela FEBRABAN
// para identificar o operação que está contida no lote
type Lancamento int

const (
	CreditoContaCorrente                     Lancamento = 1
	ChequePagamento                          Lancamento = 2
	DocTed                                   Lancamento = 3
	CartaoSalario                            Lancamento = 4 // Somente para servico PatamentoSalarios(30)
	CreditoEmContaPoupanca                   Lancamento = 5
	OpADisposicao                            Lancamento = 10
	PagamentoContasTributosComCodigoDeBarras Lancamento = 11
	TributoDarfNormal                        Lancamento = 16
	TributoGps                               Lancamento = 17
	TributoDarfSimples                       Lancamento = 18
	TributoIptu                              Lancamento = 19
	PagamentoComAutenticacao                 Lancamento = 20
	TributoDarj                              Lancamento = 21
	TributoGareSpIcms                        Lancamento = 22
	TributoGareSpDr                          Lancamento = 23
	TributoGareSpItcmd                       Lancamento = 24
	TributoIpva                              Lancamento = 25
	TributoLicenciamento                     Lancamento = 26
	TributoDpvat                             Lancamento = 27
	LiquidacaoTitulosProprioBanco            Lancamento = 30
	PagamentoTitulosOutrosBancos             Lancamento = 31
	ExtratoContaCorrente                     Lancamento = 40
	TedOutraTitularidade                     Lancamento = 41
	TedMesmaTitularidade                     Lancamento = 43
	TedTransferenciaContaInvestimento        Lancamento = 44
	DebitoContaCorrent                       Lancamento = 50
	ExtratoGestaoCaixa                       Lancamento = 70
	DepositoJudicialContaCorrente            Lancamento = 71
	DepositoJudicialPoupanca                 Lancamento = 72
	ExtratoContaInvestimento                 Lancamento = 73
)

// Inscricao trata-se do código que identifica o tipo de inscrição
// da Empresa ou Pessoa Física perante uma instituição governamental
type Inscricao int

const (
	Isento   Inscricao = 0
	Cpf      Inscricao = 1
	CgcCnpj  Inscricao = 2
	PisPasep Inscricao = 3
	Outros   Inscricao = 9
)

type Pagamento int

const (
	DebitoContaCorrent            Pagamento = 1
	DebitoEmprestimoFinanciamento Pagamento = 2
	DebitoCartaoCredito           Pagamento = 3
)
