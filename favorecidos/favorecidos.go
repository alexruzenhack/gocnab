package favorecidos

import (
	"errors"
	"fmt"
)

type Favorecido struct {
	Camara        Camara
	Banco         uint16
	ContaCorrente ContaCorrente
	Nome          string
}

func (cc Camara) Processar() string {
	return fmt.Sprintf("%03d", cc)
}

func CriarFavorecido(camara Camara, banco uint16, cc ContaCorrente, nome string) (*Favorecido, error) {
	if _, ok := Camara_chave[uint16(camara)]; !ok {
		return nil, errors.New("Código da Camara não permitido")
	}
	if camara > 999 {
		return nil, errors.New("Código da Camara deve ter no máximo 3 dígitos")
	} else if banco > 999 {
		return nil, errors.New("Código do Banco deve ter no máximo 3 dígitos")
	} else if len(nome) > 30 {
		return nil, errors.New("Nome do Favorecido deve ter no máximo 30 caracteres")
	}
	return &Favorecido{Camara: camara, Banco: banco, ContaCorrente: cc, Nome: nome}, nil
}

func (f Favorecido) Processar() string {
	return f.Camara.Processar() + fmt.Sprintf("%03d", f.Banco) + f.ContaCorrente.Processar() + fmt.Sprintf("%30s", f.Nome)
}

type Camara uint16

const (
	Camara_TED      Camara = 18
	Camara_DOC      Camara = 700
	Camara_TED_ISPB Camara = 988
)

var Camara_chave = map[uint16]string{
	18:  "TED",
	700: "DOC",
	988: "TED_ISPB",
}

var Camara_valor = map[string]uint16{
	"TED":      18,
	"DOC":      700,
	"TED_ISPB": 988,
}

type ContaCorrente struct {
	Agencia Agencia
	Conta   Conta
	Dv      string
}

func CriarContaCorrente(agencia Agencia, conta Conta) *ContaCorrente {
	sDv := conta.Dv
	if len(conta.Dv) == 2 {
		sDv = string(sDv[1]) // pega segundo dígito
	}
	return &ContaCorrente{Agencia: agencia, Conta: conta, Dv: sDv}
}

func (cc ContaCorrente) Processar() string {
	return cc.Agencia.Processar() + cc.Conta.Processar() + cc.Dv
}

type Agencia struct {
	Codigo uint32
	Dv     string
}

func CriarAgencia(codigo uint32, dv string) (*Agencia, error) {
	if codigo > 99999 {
		return nil, errors.New("Código da Agencia deve ter no máximo 5 dígitos")
	} else if len(dv) > 1 {
		return nil, errors.New("Dígito Verificador da Agencia deve ter no máximo 1 dígito")
	}
	return &Agencia{Codigo: codigo, Dv: dv}, nil
}

func (a Agencia) Processar() string {
	sCodigo := fmt.Sprintf("%05d", a.Codigo)
	return sCodigo[:5] + a.Dv[:1]
}

type Conta struct {
	Numero uint
	Dv     string
}

func CriarConta(numero uint, dv string) (*Conta, error) {
	if numero > 999999999999 {
		return nil, errors.New("Número da Conta deve ter no máximo 12 dígitos")
	} else if len(dv) > 2 {
		return nil, errors.New("Dígito Verificador da Conta deve ter no máximo 2 digitos")
	}
	return &Conta{Numero: numero, Dv: dv}, nil
}

func (c Conta) Processar() string {
	sNumero := fmt.Sprintf("%012d", c.Numero)
	return sNumero[:12] + c.Dv[:1]
}
