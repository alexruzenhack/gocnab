package endereco

import (
	"errors"
	"fmt"
)

type EnderecoEmpresa struct {
	// Texto referente a localização para entrega de correspondência.
	// Utilizado também para endereço de e-mail para entrega eletrônica
	// da informação e para número de celular para envio de mensagem SMS
	// G032
	// Tamanho: 30
	Logradouro string

	// G032
	// Tamanho: 5
	Numero uint32

	// G032
	// Tamanho: 15
	Complemento string

	// Texto referente ao nome do município componente do endereço utilizado
	// para entrega de correspondência
	// G033
	// Tamanho: 20
	Cidade string

	// Código adotado pela EBCT, para identificação de logradouros
	// G034
	// Tamanho: 5
	Cep uint32

	// Código adotado pela EBCT, para complemetanção do código de CEP
	// G035
	// Tamanho: 3
	ComplementoCep string

	// Código do estado, unidade da federação componente do endereço utilizado
	// para entrega de correspondência
	// G036
	// Tamanho: 2
	Estado string
}

func CriarEnderecoEmpresa(Logradouro string, Numero uint32, Complemento, Cidade string, Cep uint32, ComplementoCep, Estado string) (EnderecoEmpresa, error) {
	if len(Logradouro) > 30 {
		return EnderecoEmpresa{}, errors.New("Logradouro deve ter até 30 caracteres")
	} else if Numero > 99999 {
		return EnderecoEmpresa{}, errors.New("Numero deve ter até 5 dígitos")
	} else if len(Complemento) > 15 {
		return EnderecoEmpresa{}, errors.New("Complemento deve ter até 15 caracteres")
	} else if len(Cidade) > 20 {
		return EnderecoEmpresa{}, errors.New("Cidade deve ter até 20 caracteres")
	} else if Cep > 99999 {
		return EnderecoEmpresa{}, errors.New("Cep deve ter até 5 dígitos")
	} else if len(ComplementoCep) > 3 {
		return EnderecoEmpresa{}, errors.New("ComplementoCep deve ter até 3 caracteres")
	} else if len(Estado) > 2 {
		return EnderecoEmpresa{}, errors.New("Estado deve ter até 2 caracteres")
	}
	return EnderecoEmpresa{
		Logradouro,
		Numero,
		Complemento,
		Cidade,
		Cep,
		ComplementoCep,
		Estado,
	}, nil
}

func (e EnderecoEmpresa) Processar() string {
	sLogradouro := fmt.Sprintf("%-30s", e.Logradouro)
	sNumero := fmt.Sprintf("%05d", e.Numero)
	sComplemento := fmt.Sprintf("%-15s", e.Complemento)
	sCidade := fmt.Sprintf("%-20s", e.Cidade)
	sCep := fmt.Sprintf("%05d", e.Cep)
	sComplementoCep := fmt.Sprintf("%-3s", e.ComplementoCep)
	sEstado := fmt.Sprintf("%-2s", e.Estado)
	return sLogradouro[:30] + sNumero[:5] + sComplemento[:15] + sCidade[:20] + sCep[:5] + sComplementoCep[:3] + sEstado[:2]
}
