package contacorrente

import (
	"testing"
)

func TestPodeCriarContaCorrente(t *testing.T) {
	agencia := Agencia{12345, "1"}
	conta := Conta{678, "2"} // Dv de 1 dígito
	resultado := CriarContaCorrente(agencia, conta)
	if 12345 != resultado.Agencia.Codigo || "1" != resultado.Agencia.Dv {
		t.Error("Erro ao configurar a propriedade Agencia do objeto ContaCorrente")
	} else if 678 != resultado.Conta.Numero || "2" != resultado.Conta.Dv {
		t.Error("Erro ao configurar a propriedade Conta do objeto ContaCorrente")
	} else if "2" != resultado.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto ContaCorrente")
	}
}

func TestPodeCriarContaCorrenteComDvDe2Digitos(t *testing.T) {
	agencia := Agencia{67890, "2"}
	conta := Conta{901, "34"} // Dv de 2 dígitos
	resultado := CriarContaCorrente(agencia, conta)
	if "4" != resultado.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto ContaCorrente quando Dv tem dois digitos")
	}
}

func TestPodeCriarContaCorrenteComDvAcimaDe2Digitos(t *testing.T) {
	agencia := Agencia{23465, "3"}
	conta := Conta{987, "654"} // Dv acima de 2 dígitos
	resultado := CriarContaCorrente(agencia, conta)
	if "5" != resultado.Dv {
		t.Error("Erro ao configurar a propriedade Dv do objeto ContaCorrente quando Dv de Conta é maior que 2 dígitos, Dv de ContaCorrente deve ser o segundo dígito")
	}
}

func TestPodeProcessarContaCorrente(t *testing.T) {
	agencia := Agencia{123456, "7"}
	conta := Conta{891011121314, "15"}
	cc := CriarContaCorrente(agencia, conta)
	resultado := cc.Processar()
	if len(resultado) != 20 {
		t.Error("Erro no padrão de processamento da ContaCorrente")
	}
}

func TestStringDeterminadaAoProcessarContaCorrente(t *testing.T) {
	agencia := Agencia{12345, "7"}
	conta := Conta{891011121314, "15"}
	cc := CriarContaCorrente(agencia, conta)
	resultado := cc.Processar()
	if "12345789101112131415" != resultado {
		t.Error("Erro na determinação do resultado ao processar uma ContaCorrente")
	}
}

func TestPodeProcessarContaCorrenteComDvAcimaDe2Digitos(t *testing.T) {
	agencia := Agencia{123456, "7"}
	conta := Conta{891011121314, "267"}
	cc := CriarContaCorrente(agencia, conta)
	resultado := cc.Processar()
	if len(resultado) != 20 {
		t.Error("Erro no padrão de processamento da ContaCorrente com propriedade Dv além do limite de 2 dígitos")
	}
}

func TestStringDeterminadaAoProcessarContaCorrenteComDvAcimaDe2Digitos(t *testing.T) {
	agencia := Agencia{123456, "7"}
	conta := Conta{891011121314, "267"}
	cc := CriarContaCorrente(agencia, conta)
	resultado := cc.Processar()
	if "12345789101112131426" != resultado {
		t.Error("Erro na determinação do resultado ao processar uma ContaCorrente")
	}
}
