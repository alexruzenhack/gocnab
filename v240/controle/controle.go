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
	//
	// Tamanho: 3
	Banco uint16 // G001

	// Número sequencial para identificar univocamente um lote de serviço.
	// Criado e controlado pelo cliente.
	//
	// Tamanho: 4
	Lote uint16 // *G002

	// Tamanho: 1
	Registro Registro // *G003
}

func (c Controle) Processar() string {
	sBanco := fmt.Sprintf("%03d", c.Banco)
	sLote := fmt.Sprintf("%04d", c.Lote)
	return sBanco[:3] + sLote[:4] + c.Registro.Processar()
}

// Ao criar Header de Arquivo, a propriedade Lote recebe o valor padrão 0000,
// e o Registro recebe o valor padrão 0 equivalente a "Header de Arquivo"
func CriarControleHeaderArquivo(banco uint16) (Controle, error) {
	if banco > 999 {
		return Controle{}, errors.New("Código do Banco deve ter no máximo 3 dígitos")
	}
	return Controle{Banco: banco, Lote: 0, Registro: HeaderArquivo}, nil
}

// Ao criar Trailer de Arquivo, a propriedade Lote recebe o valor padrão 9999,
// e o Registro recebe o valor padrão 9 equivalente a "Trailer de Arquivo"
func CriarControleTrailerArquivo(headerArquivo Controle) Controle {
	return Controle{
		Banco:    headerArquivo.Banco,
		Lote:     9999,
		Registro: TrailerArquivo}
}

// Ao criar Header de Lote, a propriedade Registro recebe o valor padrão 1,
// mas a propriedade Lote deve iniciar em 001 e estabelecer uma sequencia
// aritmética progrssiva com razão 1
func CriarControleHeaderLote(banco uint16, lote uint16) (Controle, error) {
	if banco > 999 {
		return Controle{}, errors.New("Código do Banco deve ter no máximo 3 dígitos")
	}
	if lote == 0 {
		return Controle{}, errors.New("Lote do Arquivo deve começar em 0001")
	}
	if lote > 9999 {
		return Controle{}, errors.New("Lote do Arquivo deve ter no máximo 4 dígitos")
	}
	return Controle{Banco: banco, Lote: lote, Registro: HeaderLote}, nil
}

// Ao criar Trailer de Lote, a propriedade Registro recebe o valor padrão 5,
// e deve receber as propriedades Banco e Lote do Header de Lote que lhe deu
// origen
func CriarControleTrailerLote(headerLote Controle) Controle {
	return Controle{
		Banco:    headerLote.Banco,
		Lote:     headerLote.Lote,
		Registro: TrailerLote}
}

// Ao criar Controle Detalhe para o serviço, a propriedade Registro
// recebe o valor padrão 3 e deve receber as propriedades Banco e Lote
// do Header de Lote que lhe deu origem
func CriarControleDetalhe(header Controle) Controle {
	return Controle{
		Banco:    header.Banco,
		Lote:     header.Lote,
		Registro: Detalhe}
}

// Registro trata-se do código numérico usado pela FEBRABAN
// para identificar o tipo de registro do documento
// G003
type Registro uint16

func (r Registro) Processar() string {
	sRegistro := fmt.Sprint(r)
	return sRegistro[:1]
}

const (
	HeaderArquivo         Registro = 0
	HeaderLote            Registro = 1
	RegistrosIniciaisLote Registro = 2
	Detalhe               Registro = 3
	RegistrosFinaisLote   Registro = 4
	TrailerLote           Registro = 5
	TrailerArquivo        Registro = 9
)
