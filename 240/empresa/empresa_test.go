package empresa

import (
	pkgCc "cnab/240/contacorrente"
	"fmt"
	"strconv"
	"testing"
)

func TestPodeCriarInscricao(t *testing.T) {
	inscricao, err := CriarInscricao(Tipo_CPF, 12345678901)
	if err != nil {
		t.Error(err)
	} else if Tipo_CPF != inscricao.Tipo {
		t.Error("Erro ao configurar a propriedade Tipo")
	} else if 12345678901 != inscricao.Numero {
		t.Error("Erro ao configurar a propriedade Numero")
	}
}

func TestCriarInscricao(t *testing.T) {
	cenariosTest := []struct {
		campo  string
		tipo   Tipo
		numero uint64
		saida  string
	}{
		{
			"Tipo",
			Tipo(10),
			1,
			"Tipo não encontrado",
		},
		{
			"Numero",
			Tipo_CPF,
			123456789012340,
			"Número acima de 14 dígitos",
		},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Cria Inscricao
			_, err := CriarInscricao(cenario.tipo, cenario.numero)

			// 2. Verifica se os requisitos foram aplicados no campo
			if err == nil {
				t.Errorf("Erro ao verificar as restrições no campo %s", cenario.campo)
			}

			// 3. Verifica se mensagem de erro é apropriada
			if esperado, obtido := cenario.saida, fmt.Sprint(err); esperado != obtido {
				t.Errorf("Desejado: %s \n Obtido: %s", esperado, obtido)
			}
		})
	}
}

func TestPodeProcessarInscricao(t *testing.T) {
	inscricao := Inscricao{Tipo_CPF, 123}
	resultado := inscricao.Processar()
	if len(resultado) != 15 {
		t.Error("Erro no padrão de processamento de Inscricao")
	} else if "100000000000123" != resultado {
		t.Error("Erro na determinação do processamento de Inscricao")
	}
}

func TestProcessarInscricao(t *testing.T) {
	cenariosTest := []struct {
		tipo     Tipo
		numero   uint64
		esperado string
	}{
		{12, 123, "100000000000123"},
		{1, 123456789012340, "112345678901234"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Cria inscrição com os parametros do cenario
			inscricao := Inscricao{cenario.tipo, cenario.numero}

			// 2. Processa função
			resultado := inscricao.Processar()

			// 3. Verifica tamanho do resultado
			if esperado, obtido := 15, len(resultado); esperado != obtido {
				t.Errorf("Erro no padrão de processamento\nEsperado: %d\nObtido: %d", esperado, obtido)
				return
			}

			// 4. Verifica valor do resultado
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro na determinação do processamento\nEsperado: %s\nObtido: %s", esperado, obtido)
				return
			}
		})
	}
}

func TestPodeCriarEmpresa(t *testing.T) {
	inscricao := Inscricao{1, 123}
	ag := pkgCc.Agencia{1, "1"}
	ca := pkgCc.Conta{123, "2"}
	cc := pkgCc.ContaCorrente{ag, ca, "2"}
	empresa, err := CriarEmpresa(inscricao, "A", cc, "Minha Empresa")
	if err != nil {
		t.Error(err)
	} else if inscricao != empresa.Inscricao {
		t.Error("Erro ao configurar Inscricao")
	} else if "A" != empresa.Convenio {
		t.Error("Erro ao configurar Convenio")
	} else if cc != empresa.ContaCorrente {
		t.Error("Erro ao configurar ContaCorrente")
	} else if "Minha Empresa" != empresa.Nome {
		t.Error("Erro ao configurar Nome")
	}
}

func TestCriarEmpresa(t *testing.T) {
	cenariosTeste := []struct {
		campo                 string
		inscricao             Inscricao
		convenio, nome, saida string
	}{
		{
			"Convenio",
			Inscricao{1, 1},
			"XXXXXXXXXXXXXXXXXXXX20",
			"Nome Qualquer",
			"Erro ao tentar criar Empresa com propriedade Convenio acima de 20 dígitos",
		},
		{
			"Nome",
			Inscricao{1, 1},
			"CODIGO-CONVENIO",
			"NOME EMPRESAAAAAAAAAAAAAAAAAAA30",
			"Erro ao tentar criar Empresa com propriedade Nome acima de 30 dígitos",
		},
	}

	ccPadrao := pkgCc.ContaCorrente{}

	for i, cenario := range cenariosTeste {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Cria Empresa com campo fora do requisito
			_, err := CriarEmpresa(cenario.inscricao, cenario.convenio, ccPadrao, cenario.nome)

			// 2. Verifica se o requisito foi aplicado no campo
			if err == nil {
				t.Errorf("Erro ao verificar as restrições do campo %s", cenario.campo)
				return
			}

			// 3. Verifica mensagens de erro
			if desejado, obtido := cenario.saida, err; desejado != fmt.Sprint(obtido) {
				t.Errorf("Desejado: %v, mas foi Obtido: %s", desejado, obtido)
				return
			}
		})
	}
}

func TestPodeProcessarEmpresa(t *testing.T) {
	inscricao := Inscricao{1, 123}
	ag := pkgCc.Agencia{1, "1"}
	ca := pkgCc.Conta{123, "2"}
	cc := pkgCc.ContaCorrente{ag, ca, "2"}
	empresa := Empresa{inscricao, "A", cc, "Nome Minha Empresa"}
	resultado := empresa.Processar()
	if esperado, obtido := 85, len(resultado); esperado != obtido {
		t.Error("Erro no padrão de processamento da Empresa")
		t.Errorf("Esperado %d, mas foi obtido %d", esperado, obtido)
	} else if esperado, obtido := "100000000000123A                   00001100000000012322Nome Minha Empresa            ", resultado; esperado != obtido {
		t.Errorf("Erro na determinação do processamento da Empresa\nEsperado %s\nObtido %s", esperado, obtido)
	}
}

func TestProcessaEmpresa(t *testing.T) {
	inscricao := Inscricao{Tipo_CPF, 123}
	cc := pkgCc.ContaCorrente{pkgCc.Agencia{1, "1"}, pkgCc.Conta{123, "1"}, "1"}

	cenariosTest := []struct {
		convenio, nome, esperado string
	}{
		{"XXXXXXXXXXXXXXXXXXXX21", "NOME", inscricao.Processar() + "XXXXXXXXXXXXXXXXXXXX" + cc.Processar() + "NOME                          "},
		{"CONVENIO", "NNNNNNNNNNNNNNNNNNNNNNNNNNNNNN31", inscricao.Processar() + "CONVENIO            " + cc.Processar() + "NNNNNNNNNNNNNNNNNNNNNNNNNNNNNN"},
	}

	for i, cenario := range cenariosTest {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			// 1. Criar empresa
			empresa := Empresa{inscricao, cenario.convenio, cc, cenario.nome}

			// 2. Processar empresa
			resultado := empresa.Processar()

			// 3. Verificar tamanho do resultado
			if esperado, obtido := 85, len(resultado); esperado != obtido {
				t.Errorf("Erro no padrão de processamento da Empresa\nEsperado: %d\nObtido: %d", esperado, obtido)
				return
			}

			// 4. Verificar valor do resultado
			if esperado, obtido := cenario.esperado, resultado; esperado != obtido {
				t.Errorf("Erro na determinação do processamento da Empresa\nEsperado: %s\nObtido: %s", esperado, obtido)
				return
			}
		})
	}
}
