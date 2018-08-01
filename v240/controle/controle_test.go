package controle

import (
	"fmt"
	"testing"
)

func TestPodeCriarControleHeaderArquivo(t *testing.T) {
	cha, ok := CriarControleHeaderArquivo(1)
	if ok != nil {
		t.Error(ok)
	} else if 1 != cha.Banco {
		t.Error("Erro ao configurar a propriedade Banco")
	} else if 0 != cha.Lote {
		t.Error("Erro ao configurar a propriedade Lote")
	} else if 0 != cha.Registro {
		t.Error("Erro ao configurar a propriedade Registro")
	}
}

func TestErroAoCriarControleHeaderArquivoComBancoMaiorQue3Digitos(t *testing.T) {
	_, ok := CriarControleHeaderArquivo(1000)
	if ok == nil {
		t.Error("Erro, a propriedade Banco deve ter até 3 dígitos")
	} else if "Código do Banco deve ter no máximo 3 dígitos" != fmt.Sprint(ok) {
		t.Error("Mensagem inapropriada ao criar Controle Header Arquivo com Banco acima de 3 dígitos")
	}
}

func TestPodeCriarControleTrailerArquivo(t *testing.T) {
	cha, _ := CriarControleHeaderArquivo(100)
	cta := CriarControleTrailerArquivo(cha)
	if cha.Banco != cta.Banco {
		t.Error("Erro ao configurar a propriedade Banco")
	} else if 9999 != cta.Lote {
		t.Error("Erro ao configurar a propriedade Lote")
	} else if 9 != cta.Registro {
		t.Error("Erro ao configurar a propriedade Registro")
	}
}

// TODO: TestCriarControleHeaderLote
func TestPodeCriarControleHeaderLote(t *testing.T) {
	chl, ok := CriarControleHeaderLote(1, 1)
	if ok != nil {
		t.Error(ok)
	} else if 1 != chl.Banco {
		t.Error("Erro ao configurar a propriedade Banco")
	} else if 1 != chl.Lote {
		t.Error("Erro ao configurar a propriedade Lote")
	} else if 1 != chl.Registro {
		t.Error("Erro ao configurar a propriedade Registro")
	}
}

var tabelaErrosCriarControleHeaderLote = []struct {
	campo string
	banco uint16
	lote  uint16
	saida string
}{
	{"Banco", 1000, 1, "Código do Banco deve ter no máximo 3 dígitos"},
	{"Lote Zero", 1, 0, "Lote do Arquivo deve começar em 0001"},
	{"Lote", 1, 10000, "Lote do Arquivo deve ter no máximo 4 dígitos"},
}

func TestErroAoCriarControleHeaderLoteComDigitosAlemDoLimite(t *testing.T) {
	for _, entry := range tabelaErrosCriarControleHeaderLote {
		_, ok := CriarControleHeaderLote(entry.banco, entry.lote)
		if ok == nil {
			t.Errorf("Erro, a propriedade %s está fora do critério de aceitação", entry.campo)
		} else if entry.saida != fmt.Sprint(ok) {
			t.Errorf("Mensagem de erro inapropriada ao tentar configurar propriedade %s fora do critério permitido", entry.campo)
		}
	}
}

func TestPodeCriarControleTrailerLote(t *testing.T) {
	chl, _ := CriarControleHeaderLote(1, 1)
	ctl := CriarControleTrailerLote(chl)
	if chl.Banco != ctl.Banco {
		t.Error("Erro ao configurar propriedade Banco")
	} else if chl.Lote != ctl.Lote {
		t.Error("Erro ao configurar propriedade Lote")
	} else if 5 != ctl.Registro {
		t.Error("Erro ao configurar propriedade Registro")
	}
}

func TestPodeCriarControleDetalhe(t *testing.T) {
	chl, _ := CriarControleHeaderLote(1, 1)
	cd := CriarControleDetalhe(chl)
	if chl.Banco != cd.Banco {
		t.Error("Erro ao configurar propriedade Banco")
	} else if chl.Lote != cd.Lote {
		t.Error("Erro ao configurar propriedade Lote")
	} else if 3 != cd.Registro {
		t.Error("Erro ao configurar propriedade Detalhe")
	}
}

func TestPodeProcessarControle(t *testing.T) {
	controle := Controle{1, 1, 1}
	resultado := controle.Processar()
	if len(resultado) != 8 {
		t.Error("Erro no padrão de processamento do Controle")
	} else if "00100011" != resultado {
		t.Error("Erro na determinação do resultado ao processar Controle")
	}
}

var tableErrosProcessarControle = []struct {
	campo    string
	controle Controle
	saida    string
}{
	// Banco com mais de 3 dígitos
	{"Banco", Controle{1000, 2, 3}, "10000023"},
	// Lote com mais de 4 dígitos
	{"Lote", Controle{1, 10000, 3}, "00110003"},
	// Registro com mais de 1 dígito
	{"Registro", Controle{1, 2, 43}, "00100024"},
}

func TestPodeProcessarControleComCamposAlemDoLimite(t *testing.T) {
	for _, entry := range tableErrosProcessarControle {
		resultado := entry.controle.Processar()
		if len(resultado) != 8 {
			t.Errorf("Erro no padrão de processamento do Controle com campo %s além do limite", entry.campo)
		} else if entry.saida != resultado {
			t.Errorf("Erro na determinação do resultado ao processar Controle com campo %s além do limite", entry.campo)
		}
	}
}
