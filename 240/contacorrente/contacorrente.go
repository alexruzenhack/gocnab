package contacorrente

import (
	"errors"
	"fmt"
)

type ContaCorrente struct {
	Agencia Agencia
	Conta   Conta
	Dv      string
}

func CriarContaCorrente(agencia Agencia, conta Conta) ContaCorrente {
	sDv := conta.Dv
	if len(conta.Dv) >= 2 {
		sDv = string(sDv[1]) // pega segundo dígito
	}
	return ContaCorrente{Agencia: agencia, Conta: conta, Dv: sDv}
}

func (cc ContaCorrente) Processar() string {
	return cc.Agencia.Processar() + cc.Conta.Processar() + cc.Dv
}

type Agencia struct {
	Codigo uint32
	Dv     string
}

func CriarAgencia(codigo uint32, dv string) (Agencia, error) {
	if codigo > 99999 {
		return Agencia{}, errors.New("Código da Agencia deve ter no máximo 5 dígitos")
	} else if len(dv) > 1 {
		return Agencia{}, errors.New("Dígito Verificador da Agencia deve ter no máximo 1 dígito")
	}
	return Agencia{Codigo: codigo, Dv: dv}, nil
}

func (a Agencia) Processar() string {
	sCodigo := fmt.Sprintf("%05d", a.Codigo)
	return sCodigo[:5] + a.Dv[:1]
}

type Conta struct {
	Numero uint
	Dv     string
}

func CriarConta(numero uint, dv string) (Conta, error) {
	if numero > 999999999999 {
		return Conta{}, errors.New("Número da Conta deve ter no máximo 12 dígitos")
	} else if len(dv) > 2 {
		return Conta{}, errors.New("Dígito Verificador da Conta deve ter no máximo 2 digitos")
	}
	return Conta{Numero: numero, Dv: dv}, nil
}

func (c Conta) Processar() string {
	sNumero := fmt.Sprintf("%012d", c.Numero)
	return sNumero[:12] + c.Dv[:1]
}
