package arquivo

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

const vLayoutArquivo uint16 = 103

type Arquivo struct {
	// Código adotado pela FEBRABAN para qualificar o envio ou devolução de arquivo
	// entre a Empresa Cliente e o Banco prestador dos Serviços
	// G015
	// Tamanho: 1
	Codigo Codigo

	// Data de criação do arquivo
	// Deve seguir o formato DDMMAAAA
	// G016
	// Tamanho: 8
	DataGeracao uint32

	// Hora de criação do arquivo
	// Deve seguir o formato HHMMSS
	// G017
	// Tamanho: 6
	HoraGeracao uint32

	// Número sequencial adotado e controlado pelo responsável pela geração do arquivo
	// para ordenar a disposição dos arquivos encaminhados
	// Evoluir uma unidade a cada header de arquivo
	// G018*
	// Tamanho: 6
	Sequencia uint32

	// Código adotado pela FEBRABAN para identificar qual a versão de layout do arquivo
	// encaminhado
	// Composto de "Versão (2 dígitos)" e "Release (1 dígito)"
	// Assume como valor padrão a versão mais atual refletido na constante vLayouArquivo
	// G019*
	// Tamaho: 3
	LayoutArquivo uint16

	// Densidade de gravação (BPI), do arquivo encaminhado
	// G020
	// Tamanho: 5
	Densidade Densidade
}

type Codigo uint8

// Domínio do Código
const (
	Codigo_REMESSA Codigo = 1
	Codigo_RETORNO Codigo = 2
)

type Densidade uint32

const (
	Densidade_1600_BPI Densidade = 1600
	Densidade_6250_BPI Densidade = 6250
)

var Densidade_Valor = map[uint32]string{
	1600: "1600_BPI",
	6250: "6250_BPI",
}

func CriarArquivo(seq uint32, dsde Densidade) (Arquivo, error) {
	if seq == 0 {
		return Arquivo{}, errors.New("Sequencia deve ser maior do que zero")
	} else if seq > 999999 {
		return Arquivo{}, errors.New("Sequencia deve ter até 6 dígitos")
	} else if _, ok := Densidade_Valor[uint32(dsde)]; !ok {
		return Arquivo{}, errors.New("Valor de Densidade não permitido")
	}

	now := time.Now()
	data, err := strconv.Atoi(now.Format("02012006"))
	if err != nil {
		panic(err)
	}
	hora, err := strconv.Atoi(now.Format("150405"))
	if err != nil {
		panic(err)
	}

	return Arquivo{
		Codigo_REMESSA,
		uint32(data),
		uint32(hora),
		seq,
		vLayoutArquivo,
		dsde,
	}, nil
}

func (c Codigo) processar() string {
	sCodigo := fmt.Sprint(c)
	return sCodigo[:1]
}

func (d Densidade) processar() string {
	sDensidade := fmt.Sprintf("%05d", d)
	return sDensidade[:5]
}

func (a Arquivo) Processar() string {
	sData := fmt.Sprintf("%08d", a.DataGeracao)
	sHora := fmt.Sprintf("%06d", a.HoraGeracao)
	sSeq := fmt.Sprintf("%06d", a.Sequencia)
	sLayout := fmt.Sprintf("%03d", a.LayoutArquivo)
	return a.Codigo.processar() + sData[:8] + sHora[:6] + sSeq[:6] + sLayout[:3] + a.Densidade.processar()
}
