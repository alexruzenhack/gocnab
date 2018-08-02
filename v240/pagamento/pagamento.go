package pagamento

import (
	"cnab/v240/controle"
	"cnab/v240/empresa"
	"cnab/v240/servico"
)

type HeaderLote struct {
	Controle controle.Controle

	Servico servico.ServicoHeader

	// valor padrão caracter vazio ' '
	// G004
	// tamanho: 1
	CnabHeader string

	Empresa empresa.Empresa

	// Texto referente a mensagens que serão impressas nos documentos
	// e/ou avisos a serem emitidos
	// Informação 1: Mensagem genérica
	// *G031
	// tamanho: 40
	Mensagem string

	// Endereço da Empresa
	// Texto referente a localização da rua/avenida, número, complemento,
	// e bairro utilizado para entrega de correspondencia. Utilizado
	// também para endereço de e-mail para entrega eletrônica de informação
	// e para número de celular para envio de mensagem SMS
	// G032
	// tamanho: 30
	Logradouro string

	// G032
	// tamanho: 5
	Numero int

	// G032
	// tamanho: 15
	Complemento string

	// Texto referente ao nome do município componente do endereço
	// utilizado para entrega de correspondencia
	// G033
	// tamanho: 20
	Cidade string

	// Código adotado pela EBCT (Empresa Brasileira de Correios e Telégrafos)
	// para identificação de logradouros
	// G034
	// tamanho: 5
	Cep int

	// Código para complementação do código do CEP
	// G035
	// tamanho: 3
	ComplementoCep string

	// Código do estado, unidade da federação componente do endereço
	// utilizado para entrega de correspondência
	// G036
	// tamanho: 2
	Estado string

	// Possibilitar ao Pagador, mediante acordo com seu Banco de Relacionamento
	// a forma de pagamento do compromisso
	// P014
	// tamanho: 2
	FormaPagamento Pagamento

	// Cnab
	// G004
	CnabTrailer int

	// Código adotado pela FEBRABAN para identificar as ocorrências
	// detectadas no processamento
	// Pode-se informar até 5 ocorrências simultaneamente, cada uma delas
	// codificada com dois dígitos
	// *G059
	// tamanho: 10
	Ocorrencias int
}
