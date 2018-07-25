package favorecidos

import (
	"fmt"
	"testing"
)

func TestPodeCriarConta(t *testing.T) {
	result, ok := CriarConta(3244, "1")
	if ok != nil {
		t.Error(ok)
	} else if 3244 != result.Numero {
		t.Error("Erro ao configurar a propriedade Numero do objeto Conta")
	} else if "1" != result.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto Conta")
	}
}

func TestPodeCriarContaComDvDeDoisDigitos(t *testing.T) {
	result, ok := CriarConta(9854, "43")
	if ok != nil {
		t.Error(ok)
	} else if "43" != result.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto Conta")
	}
}

func TestErroAoCriarContaComNumeroMaiorQue12Digitos(t *testing.T) {
	_, ok := CriarConta(9876543210013, "0")
	if ok == nil {
		t.Error("Erro, a propriedade Numero do objeto Conta deve ter até 12 dígitos")
	}
}

func TestErroAoCriarContaComDvMaiorQue2Digitos(t *testing.T) {
	_, ok := CriarConta(001, "123")
	if ok == nil {
		t.Error("Erro, a propriedade Dv do objeto Conta deve ter até 2 dígitos")
	}
}

func TestPodeProcessarConta(t *testing.T) {
	conta := &Conta{Numero: 123, Dv: "1"}
	result := conta.Processar()
	if len(result) != 13 {
		t.Error("Erro no padrão de processamento da Conta")
	}
}

func TestStringDeterminadaAoProcessarConta(t *testing.T) {
	conta := &Conta{Numero: 321, Dv: "2"}
	result := conta.Processar()
	if "0000000003212" != result {
		t.Error("Erro na determinação do resultado ao processar uma Conta")
	}
}

func TestPodeProcessarContaComDigitoVerificadorDeDoisDigitos(t *testing.T) {
	conta := &Conta{Numero: 456, Dv: "78"}
	result := conta.Processar()
	if len(result) != 13 {
		t.Error("Erro no padrão de processamento da Conta com Dígito Verificador de dois dígitos")
	}
}

func TestStringDeterminadaAoProcessarContaComDvDeDoisDigitos(t *testing.T) {
	conta := &Conta{Numero: 789, Dv: "56"}
	result := conta.Processar()
	if "0000000007895" != result {
		t.Error("Erro na determinação do resultado ao processar uma Conta com Dígito Verificador de dois dígitos")
	}
}

func TestPodeProcessarContaComDigitosAlemDoLimite(t *testing.T) {
	conta := &Conta{Numero: 1234567890013, Dv: "123"}
	result := conta.Processar()
	if len(result) != 13 {
		fmt.Print(result)
		t.Error("Erro no processamento da Conta com dígitos das propriedades Numero e Dv acima do padrão")
	}
}
