package favorecidos

import (
	"fmt"
	"testing"
)

func TestPodeCriarAgencia(t *testing.T) {
	resultado, ok := CriarAgencia(123, "0")
	if ok != nil {
		t.Error(ok)
	} else if 123 != resultado.Codigo {
		t.Error("Erro ao configurar a propriedade Codigo do objeto Agencia")
	} else if "0" != resultado.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto Agencia")
	}
}

func TestErroAoCriarAgenciaComCodigoMaiorQue5Digitos(t *testing.T) {
	_, ok := CriarAgencia(123456, "0")
	if ok == nil {
		t.Error("Erro, a propriedade Codigo do objeto Agencia deve ter até 5 dígitos")
	}
}

func TestMensagemErroAoCriarAgenciaComCodigoMaiorQue5Digitos(t *testing.T) {
	_, ok := CriarAgencia(654321, "0")
	if "Código da Agencia deve ter no máximo 5 dígitos" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Agencia com Codigo acima de 5 dígitos")
	}
}

func TestErroAoCriarAgenciaComDvMaiorQue1Digito(t *testing.T) {
	_, ok := CriarAgencia(456, "12")
	if ok == nil {
		t.Error("Erro, a propriedade Dv do objeto Agencia deve ter apenas 1 digito")
	}
}

func TestMensagemErroAoCriarAgenciaComDvMaiorQue1Digito(t *testing.T) {
	_, ok := CriarAgencia(123, "21")
	if "Dígito Verificador da Agencia deve ter no máximo 1 dígito" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Agencia com Dv acima de 1 digito")
	}
}

func TestPodeProcessarAgencia(t *testing.T) {
	agencia := &Agencia{Codigo: 123, Dv: "0"}
	resultado := agencia.Processar()
	if len(resultado) != 6 {
		t.Error("Erro no padrão de processamento da Agencia")
	}
}

func TestStringDeterminadaAoProcessarAgencia(t *testing.T) {
	agencia := &Agencia{Codigo: 456, Dv: "1"}
	resultado := agencia.Processar()
	if "004561" != resultado {
		t.Error("Erro no resultado ao processar uma Agencia")
	}
}

func TestPodeProcessarAgenciaComDigitosAlemDoLimite(t *testing.T) {
	agencia := &Agencia{Codigo: 987654, Dv: "13"}
	resultado := agencia.Processar()
	if len(resultado) != 6 || "987651" != resultado {
		t.Error("Erro no processamento da Agencia com dígitos das propriedades Codigo e Dv acima do padrão")
	}
}
