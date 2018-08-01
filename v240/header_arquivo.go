package v240

import (
	"cnab/v240/arquivo"
	"cnab/v240/controle"
	"cnab/v240/empresa"
	"errors"
	"fmt"
)

type HeaderArquivo struct {
	// G001, *G002, *G003
	// Tamanho: 8
	Controle controle.Controle

	// Texto de observações destinado para uso exclusivo da FEBRABAN
	// Deve ser preenchido com brancos
	// G004
	// Tamaho: 9
	CnabHeader string

	// *G005, *G006, *G007, *G008, *G009, *G010, *G011, *G012, G013
	// Tamanho: 85
	Empresa empresa.Empresa

	// Nome que identifica o Banco que está recebendo ou enviando o arquivo
	// G014
	// Tamanho: 30
	NomeBanco string

	// G004
	// Tamanho: 10
	CnabBody string

	// G015, G016, G017, *G018, *G019, G020
	// Tamanho: 29
	Arquivo arquivo.Arquivo

	// Texto de observações destinado para uso exclusivo do Banco
	// G021
	// Tamaho: 20
	ReservadoBanco string

	// Texto de observações destinado para uso exclusivo da Empresa
	// G022
	// Tamanho: 20
	ReservadoEmpresa string

	// G004
	// Tamanho: 29
	CnabTrailer string
}

func CriarHeaderArquivo(controle controle.Controle, empresa empresa.Empresa, arquivo arquivo.Arquivo, nomeBanco, reservadoEmpresa string) (HeaderArquivo, error) {
	if len(nomeBanco) > 30 {
		return HeaderArquivo{}, errors.New("NomeBanco deve ter até 30 dígitos")
	} else if len(reservadoEmpresa) > 20 {
		return HeaderArquivo{}, errors.New("ReservadoEmpresa deve ter até 20 dígitos")
	}

	cnabHeader := fmt.Sprintf("%-9s", "")
	cnabBody := fmt.Sprintf("%-10s", "")
	cnabTrailer := fmt.Sprintf("%-29s", "")
	reservadoBanco := fmt.Sprintf("%-20s", "")

	return HeaderArquivo{
		controle,
		cnabHeader,
		empresa,
		nomeBanco,
		cnabBody,
		arquivo,
		reservadoBanco,
		reservadoEmpresa,
		cnabTrailer}, nil
}

func (ha HeaderArquivo) Processar() string {
	sCnabHeader := fmt.Sprintf("%-9s", ha.CnabHeader)
	sCnabBody := fmt.Sprintf("%-10s", ha.CnabBody)
	sCnabTrailer := fmt.Sprintf("%-29s", ha.CnabTrailer)
	sNomeBanco := fmt.Sprintf("%-30s", ha.NomeBanco)
	sReservadoBanco := fmt.Sprintf("%-20s", ha.ReservadoBanco)
	sReservadoEmpresa := fmt.Sprintf("%-20s", ha.ReservadoEmpresa)
	return ha.Controle.Processar() + sCnabHeader[:9] + ha.Empresa.Processar() + sNomeBanco[:30] + sCnabBody[:10] + ha.Arquivo.Processar() + sReservadoBanco[:20] + sReservadoEmpresa[:20] + sCnabTrailer[:29]
}
