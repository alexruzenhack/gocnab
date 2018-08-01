package v240

import (
	ctrl "cnab/v240/controle"
	svco "cnab/v240/servico"
)

// Header
type Header struct {
	Controle ctrl.Controle

	// Servico
	// valor padrão igual a LancamentoCredito ('C')
	// tamanho: 1
	Operacao svco.Operacao // *G028

	// Quando o serviço adotado é Interoperabilidade entre Contas (23)
	// é obrigatório o preenchimento do campo 18.3C
	// tamanho: 2
	Servico svco.Servico // *G025

	// valor padrão igual a 046
	// tamanho: 2
	FormaLancamento svco.FormaLancamento // *G029

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
