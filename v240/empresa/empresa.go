package empresa

import (
	pkgCc "cnab/v240/contacorrente"
	"errors"
	"fmt"
)

type Empresa struct {
	// Tamanho: 15
	Inscricao Inscricao

	// Código adotado pelo banco para identificar o contrato entre este
	// e a empresa Cliente
	// Tamanho: 20
	Convenio string

	// Tamanho: 20
	ContaCorrente pkgCc.ContaCorrente

	// Nome que identifica a pessoa, física ou jurídica, a qual se quer
	// fazer referência
	// Tamanho: 30
	Nome string
}

type Inscricao struct {
	// Código que identifica o tipo de inscrição da Empresa ou Pessoa Física
	// perante uma instituição Governamental
	// Tamanho: 1
	Tipo Tipo

	// Número de inscrição da Empresa ou Pessoa Física perante uma instituição Governamental
	// Tamanho: 14
	Numero uint64
}

type Tipo uint8

// Domínio do tipo de inscrição
const (
	Tipo_ISENTO    Tipo = 0
	Tipo_CPF       Tipo = 1
	Tipo_CGC_CNPJ  Tipo = 2
	Tipo_PIS_PASEP Tipo = 3
	Tipo_OUTROS    Tipo = 4
)

var Tipo_Valor = map[uint8]string{
	0: "Tipo_ISENTO",
	1: "Tipo_CPF",
	2: "Tipo_CGC_CNPJ",
	3: "Tipo_PIS_PASEP",
	4: "Tipo_OUTROS",
}

func CriarInscricao(tipo Tipo, numero uint64) (Inscricao, error) {
	if _, ok := Tipo_Valor[uint8(tipo)]; !ok {
		return Inscricao{}, errors.New("Tipo não encontrado")
	} else if numero > 99999999999999 {
		return Inscricao{}, errors.New("Número acima de 14 dígitos")
	}
	return Inscricao{tipo, numero}, nil
}

func CriarEmpresa(inscricao Inscricao, convenio string, cc pkgCc.ContaCorrente, nome string) (Empresa, error) {
	if len(convenio) > 20 {
		return Empresa{}, errors.New("Erro ao tentar criar Empresa com propriedade Convenio acima de 20 dígitos")
	} else if len(nome) > 30 {
		return Empresa{}, errors.New("Erro ao tentar criar Empresa com propriedade Nome acima de 30 dígitos")
	}
	return Empresa{inscricao, convenio, cc, nome}, nil
}

func (i Inscricao) Processar() string {
	sTipo := fmt.Sprint(i.Tipo)
	sNumero := fmt.Sprintf("%014d", i.Numero)
	return sTipo[:1] + sNumero[:14]
}

func (e Empresa) Processar() string {
	sConvenio := fmt.Sprintf("%-20s", e.Convenio)
	sNome := fmt.Sprintf("%-30s", e.Nome)
	return e.Inscricao.Processar() + sConvenio[:20] + e.ContaCorrente.Processar() + sNome[:30]
}
