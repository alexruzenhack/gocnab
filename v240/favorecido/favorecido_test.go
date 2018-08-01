package favorecido

import (
	"fmt"
	"testing"
)

func TestPodeCriarFavorecido(t *testing.T) {
	ag := Agencia{123, "4"}
	c := Conta{5678, "9"}
	cc := CriarContaCorrente(ag, c)
	resultado, ok := CriarFavorecido(Camara_DOC, 1, cc, "ALEX ARAUJO RUZENHACK")
	if ok != nil {
		t.Error(ok)
	} else if 700 != uint(resultado.Camara) {
		t.Error("Erro ao configurar a propriedade Camara do objeto Favorecido")
	} else if 001 != resultado.Banco {
		t.Error("Erro ao configurar a propriedade Banco do objeto Favorecido")
	} else if "ALEX ARAUJO RUZENHACK" != resultado.Nome {
		t.Error("Erro ao configurar a propriedade Nome do objeto Favorecido")
	}
}

func TestErroAoCriarFavorecidoValorDeCamaraNaoListado(t *testing.T) {
	ag := Agencia{123, "4"}
	c := Conta{5678, "9"}
	cc := CriarContaCorrente(ag, c)
	_, ok := CriarFavorecido(Camara(1000), 1, cc, "ALEX ARAUJO RUZENHACK")
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
	_, ok := CriarFavorecido(Camara_DOC, 1, cc, "ALEX ARAUJO RUZENHACK00000000031")
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
	favorecido := Favorecido{Camara_DOC, 1, cc, "ALEX ARAUJO RUZENHACK"}
	resultado := favorecido.Processar()
	if len(resultado) != 56 {
		t.Error("Erro no padrão de processamento de Favorecido")
	} else if "70000112345678910111213141ALEX ARAUJO RUZENHACK         " != resultado {
		t.Error("Erro na determinação do resultado ao processar um Favorecido")
	}
}

var cc = ContaCorrente{
	Agencia{12345, "6"},
	Conta{789101112131, "41"},
	"1",
}

var tabelaFavorecidos = []struct {
	campo   string
	entrada Favorecido
	saida   string
}{
	// Camara com mais de 3 dígitos
	{"Camara", Favorecido{Camara(1000), 2, cc, "ALEX ARAUJO RUZENHACK"}, "10000212345678910111213141ALEX ARAUJO RUZENHACK         "},
	// Banco com mais de 3 dígitos
	{"Banco", Favorecido{Camara(100), 9999, cc, "ALEX ARAUJO RUZENHACK"}, "10099912345678910111213141ALEX ARAUJO RUZENHACK         "},
	// Nome com mais de 30 dígitos
	{"Nome", Favorecido{Camara(100), 2, cc, "ALEX ARAUJO RUZENHACK00000000031"}, "10000212345678910111213141ALEX ARAUJO RUZENHACK000000000"},
}

func TestPodeProcessarFavorecidoComCamaraAlemDoLimite(t *testing.T) {
	for _, entry := range tabelaFavorecidos {
		resultado := entry.entrada.Processar()
		if len(resultado) != 56 {
			t.Errorf("Erro no padrão de processamento de Favorecido quando %s tem dígitos além do limite permitido", entry.campo)
		} else if entry.saida != resultado {
			t.Errorf("Erro na determinação do resultado ao processar um Favorecido quando %s tem dígitos além do limite", entry.campo)
		}
	}
}
