package controle

import (
	"errors"
	"fmt"
)

type Controle struct {
	// Verificar se o código informado consta na lista de instituições
	// registradas no Banco Central
	//
	// Caso o código do banco seja equivalente ao valor 988,
	// verificar se o campo 26.3B do documento está devidamente preenchido.
	Banco uint16 // G001

	// Quando o lote está na header, então deve assumir o valor 0000
	// Quando o lete está no trailer, então deve assumir o valor 9999
	// Quando o lote representa um serviço, então deve assumir uma sequencia
	// começando por 0001 e ser acrescido de 1 unidade por serviço
	Lote uint16 // *G002

	// valor padrão igual a HeaderLote (1)
	Registro Registro // *G003
}

func CriarControleHeaderArquivo(banco uint16) (*Controle, error) {
	if banco > 999 {
		return nil, errors.New("Código do Banco deve ter no máximo 3 dígitos")
	}
	return &Controle{Banco: banco, Lote: 0, Registro: HeaderArquivo}, nil
}

func CriarControleTrailerArquivo(headerArquivo Controle) *Controle {
	return &Controle{
		Banco:    headerArquivo.Banco,
		Lote:     9999,
		Registro: TrailerArquivo}
}

func CriarControleHeaderLote(banco uint16, lote uint16) (*Controle, error) {
	if banco > 999 {
		return nil, errors.New("Código do Banco deve ter no máximo 3 dígitos")
	}
	if lote == 0 {
		return nil, errors.New("Lote do Arquivo deve começar em 0001")
	}
	if lote > 9999 {
		return nil, errors.New("Lote do Arquivo deve ter no máximo 4 dígitos")
	}
	return &Controle{Banco: banco, Lote: lote, Registro: HeaderLote}, nil
}

func CriarControleTrailerLote(headerLote Controle) *Controle {
	return &Controle{
		Banco:    headerLote.Banco,
		Lote:     headerLote.Lote,
		Registro: TrailerLote}
}

func CriarControleDetalhe(header Controle) *Controle {
	return &Controle{
		Banco:    header.Banco,
		Lote:     header.Lote,
		Registro: Detalhe}
}

func (c Controle) Processar() string {
	return fmt.Sprintf("%03d", c.Banco) + fmt.Sprintf("%04d", c.Lote) + c.Registro.Processar()
}

// Registro trata-se do código numérico usado pela FEBRABAN
// para identificar o tipo de registro do documento
// G003
type Registro uint16

const (
	HeaderArquivo         Registro = 0
	HeaderLote            Registro = 1
	RegistrosIniciaisLote Registro = 2
	Detalhe               Registro = 3
	RegistrosFinaisLote   Registro = 4
	TrailerLote           Registro = 5
	TrailerArquivo        Registro = 9
)

func (r Registro) Processar() string {
	return fmt.Sprint(r)
}
