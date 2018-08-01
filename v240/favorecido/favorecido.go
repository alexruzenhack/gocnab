package favorecido

import (
	cc "cnab/v240/contacorrente"
	"errors"
	"fmt"
)

type Favorecido struct {
	Camara        Camara
	Banco         uint16
	ContaCorrente cc.ContaCorrente
	Nome          string
}

func CriarFavorecido(camara Camara, banco uint16, cc cc.ContaCorrente, nome string) (Favorecido, error) {
	if _, ok := Camara_chave[uint16(camara)]; !ok {
		return Favorecido{}, errors.New("Código da Camara não permitido")
	}
	if camara > 999 {
		return Favorecido{}, errors.New("Código da Camara deve ter no máximo 3 dígitos")
	} else if banco > 999 {
		return Favorecido{}, errors.New("Código do Banco deve ter no máximo 3 dígitos")
	} else if len(nome) > 30 {
		return Favorecido{}, errors.New("Nome do Favorecido deve ter no máximo 30 caracteres")
	}
	return Favorecido{Camara: camara, Banco: banco, ContaCorrente: cc, Nome: nome}, nil
}

func (f Favorecido) Processar() string {
	sBanco := fmt.Sprintf("%03d", f.Banco)
	sNome := fmt.Sprintf("%-30s", f.Nome)
	return f.Camara.Processar() + sBanco[:3] + f.ContaCorrente.Processar() + sNome[:30]
}

type Camara uint16

func (c Camara) Processar() string {
	sCamara := fmt.Sprintf("%03d", c)
	return sCamara[:3]
}

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
