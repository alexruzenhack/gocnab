package totais

import (
	"errors"
	"fmt"
)

type TotaisTrailerArquivo struct {
	// Número obtido pela contagem dos lotes enviados no arquivo.
	// Somatória dos registros do tipo 1.
	// G049
	// Tamanho: 6
	QtdLotes uint32

	// Número obtido pela contagem dos registros enviados no arquivo.
	// Somatória dos registros do tipo 0, 1, 3, 5 e 9
	// G056
	// Tamanho: 6
	QtdRegistros uint32

	// Número indicativo de lotes de Conciliação Bancária enviados no arquivo.
	// Somatória dos registros do tipo 1 e Tipo de Operação = 'E'
	// Campo específico para serviço de conciliação bancária
	// *G037
	// Tamanho: 6
	QtdContasConcil uint32
}

func CriarTotaisTrailerArquivo(qtdLotes, qtdRegistros, qtdContasConcil uint32) (TotaisTrailerArquivo, error) {
	if qtdLotes > 999999 {
		return TotaisTrailerArquivo{}, errors.New("QtdLotes deve ter até 6 dígitos")
	} else if qtdRegistros > 999999 {
		return TotaisTrailerArquivo{}, errors.New("QtdRegistros deve ter até 6 dígitos")
	} else if qtdContasConcil > 999999 {
		return TotaisTrailerArquivo{}, errors.New("QtdContasConcil deve ter até 6 dígitos")
	}
	return TotaisTrailerArquivo{qtdLotes, qtdRegistros, qtdContasConcil}, nil
}

func (t TotaisTrailerArquivo) Processar() string {
	sQtdLotes := fmt.Sprintf("%06d", t.QtdLotes)
	sQtdRegistros := fmt.Sprintf("%06d", t.QtdRegistros)
	sQtdContasConcil := fmt.Sprintf("%06d", t.QtdContasConcil)
	return sQtdLotes[:6] + sQtdRegistros[:6] + sQtdContasConcil[:6]
}
