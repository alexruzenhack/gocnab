package favorecidos

import (
	"fmt"
	"testing"
)

func TestPodeCriarConta(t *testing.T) {
	resultado, ok := CriarConta(3244, "1")
	if ok != nil {
		t.Error(ok)
	} else if 3244 != resultado.Numero {
		t.Error("Erro ao configurar a propriedade Numero do objeto Conta")
	} else if "1" != resultado.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto Conta")
	}
}

func TestPodeCriarContaComDvDeDoisDigitos(t *testing.T) {
	resultado, ok := CriarConta(9854, "43")
	if ok != nil {
		t.Error(ok)
	} else if "43" != resultado.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto Conta")
	}
}

func TestErroAoCriarContaComNumeroMaiorQue12Digitos(t *testing.T) {
	_, ok := CriarConta(9876543210013, "0")
	if ok == nil {
		t.Error("Erro, a propriedade Numero do objeto Conta deve ter até 12 dígitos")
	}
}

func TestMensagemErroAoCriarContaComNumeroMaiorQue12Digitos(t *testing.T) {
	_, ok := CriarConta(9876543210013, "0")
	if "Número da Conta deve ter no máximo 12 dígitos" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Conta com Numero acima de 12 dígitos")
	}
}

func TestErroAoCriarContaComDvMaiorQue2Digitos(t *testing.T) {
	_, ok := CriarConta(001, "123")
	if ok == nil {
		t.Error("Erro, a propriedade Dv do objeto Conta deve ter até 2 dígitos")
	}
}

func TestMesagemErroAoCriarContaComDvMaiorQue2Digitos(t *testing.T) {
	_, ok := CriarConta(001, "456")
	if "Dígito Verificador da Conta deve ter no máximo 2 digitos" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Conta com Dv acima de 2 dígitos")
	}
}

func TestPodeProcessarConta(t *testing.T) {
	conta := &Conta{Numero: 123, Dv: "1"}
	resultado := conta.Processar()
	if len(resultado) != 13 {
		t.Error("Erro no padrão de processamento da Conta")
	}
}

func TestStringDeterminadaAoProcessarConta(t *testing.T) {
	conta := &Conta{Numero: 321, Dv: "2"}
	resultado := conta.Processar()
	if "0000000003212" != resultado {
		t.Error("Erro na determinação do resultado ao processar uma Conta")
	}
}

func TestPodeProcessarContaComDigitoVerificadorDeDoisDigitos(t *testing.T) {
	conta := &Conta{Numero: 456, Dv: "78"}
	resultado := conta.Processar()
	if len(resultado) != 13 {
		t.Error("Erro no padrão de processamento da Conta com Dígito Verificador de dois dígitos")
	}
}

func TestStringDeterminadaAoProcessarContaComDvDeDoisDigitos(t *testing.T) {
	conta := &Conta{Numero: 789, Dv: "56"}
	resultado := conta.Processar()
	if "0000000007895" != resultado {
		t.Error("Erro na determinação do resultadoado ao processar uma Conta com Dígito Verificador de dois dígitos")
	}
}

func TestPodeProcessarContaComDigitosAlemDoLimite(t *testing.T) {
	conta := &Conta{Numero: 1234567890013, Dv: "123"}
	resultado := conta.Processar()
	if len(resultado) != 13 || "1234567890011" != resultado {
		t.Error("Erro no processamento da Conta com dígitos das propriedades Numero e Dv acima do padrão")
	}
}
