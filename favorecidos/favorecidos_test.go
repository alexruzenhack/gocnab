package favorecidos

import (
	"fmt"
	"testing"
)

func TestPodeCriarFavorecido(t *testing.T) {
	ag := Agencia{123, "4"}
	c := Conta{5678, "9"}
	cc := CriarContaCorrente(ag, c)
	resultado, ok := CriarFavorecido(Camara_DOC, 001, cc, "ALEX ARAUJO RUZENHACK")
	if ok != nil {
		t.Error(ok)
	} else if 700 != uint(resultado.Camara) {
		t.Error("Erro ao configurar a propriedade Camara do objeto Favorecido")
	} else if 001 != resultado.Banco {
		t.Error("Erro ao configurar a propriedade Banco do objeto Favorecido")
	} else if "ALEX ARAUJO RUSENHACK" != resultado.Nome {
		t.Error("Erro ao configurar a propriedade Nome do objeto Favorecido")
	}
}

func TestErroAoCriarFavorecidoValorDeCamaraNaoListado(t *testing.T) {
	ag := Agencia{123, "4"}
	c := Conta{5678, "9"}
	cc := CriarContaCorrente(ag, c)
	_, ok := CriarFavorecido(Camara(1000), 001, cc, "ALEX ARAUJO RUZENHACK")
	if ok == nil {
		t.Error("Erro, a propriedade Camara do objeto Favorecido deve estar listada como constante")
	} else if "Código da Camara não permitido" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Favorecido com Camara não listada como constante")
	}
}

func TestErroAoCriarFavorecidoComBancoMaiorQue3Digitos(t *testing.T) {
	ag := Agencia{123, "4"}
	c := Conta{5678, "9"}
	cc := CriarContaCorrente(ag, c)
	_, ok := CriarFavorecido(Camara_DOC, 1000, cc, "ALEX ARAUJO RUZENHACK")
	if ok == nil {
		t.Error("Erro, a propriedade Banco do objeto Favorecido deve ter no máximo 3 dígitos")
	} else if "Código do Banco deve ter no máximo 3 dígitos" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Favorecido com Banco acima de 3 dígitos")
	}
}

func TestErroAoCriarFavorecidoComNomeMaiorQue30Digitos(t *testing.T) {
	ag := Agencia{123, "4"}
	c := Conta{5678, "9"}
	cc := CriarContaCorrente(ag, c)
	_, ok := CriarFavorecido(Camara_DOC, 001, cc, "ALEX ARAUJO RUZENHACK00000000031")
	if ok == nil {
		t.Error("Erro, a propriedade Nome do objeto Favorecido deve ter no máximo 30 digitos")
	} else if "Nome do Favorecido deve ter no máximo 30 caracteres" != fmt.Sprint(ok) {
		t.Error("Mensagem de erro inapropriada ao criar Favorecido com Nome acima de 30 digitos")
	}
}

func TestPodeProcessarFavorecido(t *testing.T) {
	ag := Agencia{12345, "6"}
	c := Conta{789101112131, "41"}
	cc := CriarContaCorrente(ag, c)
	favorecido := Favorecido{Camara_DOC, 001, cc, "ALEX ARAUJO RUZENHACK"}
	resultado := favorecido.Processar()
	if len(resultado) != 56 {
		t.Error("Erro no padrão de processamento de Favorecido")
	} else if "70000112345678910111213141ALEX ARAUJO RUZENHACK         " != resultado {
		fmt.Println(resultado)
		t.Error("Erro na determinação do resultado ao processar um Favorecido")
	}
}
