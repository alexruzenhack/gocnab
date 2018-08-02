package v240

import (
	"cnab/v240/controle"
	"cnab/v240/totais"
	"errors"
	"fmt"
)

type TrailerArquivo struct {
	Controle    controle.Controle
	CnabHeader  string
	Totais      totais.TotaisTrailerArquivo
	CnabTrailer string
}

func CriarTrailerArquivo(ctrle controle.Controle, ttais totais.TotaisTrailerArquivo) (TrailerArquivo, error) {
	if ctrle.Registro != controle.TrailerArquivo {
		return TrailerArquivo{}, errors.New("Controle deve possuir o Registro 9 - TrailerArquivo")
	}
	cnabHeader := fmt.Sprintf("%-9s", "")
	cnabTrailer := fmt.Sprintf("%-205s", "")
	return TrailerArquivo{ctrle, cnabHeader, ttais, cnabTrailer}, nil
}

func (ta TrailerArquivo) Processar() string {
	sCnabHeader := fmt.Sprintf("%-9s", "")
	sCnabTrailer := fmt.Sprintf("%-205s", "")
	return ta.Controle.Processar() + sCnabHeader[:9] + ta.Totais.Processar() + sCnabTrailer[:205]
}
