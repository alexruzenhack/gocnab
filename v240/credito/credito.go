package credito

import (
	"errors"
	"fmt"
)

type Credito struct {
	// Número atribuído pela Empresa (pagador) para identificar o documento
	// de pagamento
	// G064
	// Tamanho: 20
	NumeroEmpresa string

	// Data do pagamento do compromisso
	// Formato DDMMAAAA
	// P009
	// Tamanho: 8
	DataPagamento uint32

	// *G040, G041
	// Tamanho: 18
	Moeda Moeda

	// Valor do pagamento, expresso em valor corrente
	// P010
	// Tamanho: 15
	ValorPagamento uint64

	// Número atribuído pelo banco para identificar o lançamento, que será
	// utilizado nas manutenções do mesmo
	// *G043
	// Tamanho: 20
	NumeroBanco string

	// Data de efetivação do pagamento
	// A ser preenchido quando arquivo for de retorno (Código=2 no Header de Arquivo)
	// e referir-se a uma confirmação de lançamento
	// Formato DDMMAAAA
	// P003
	// Tamanho: 8
	DataReal uint32

	// Valor de efetivação do pagamento, expresso em moeda corrente
	// A ser preenchido quando arquivo for de retorno (Código=2 no Header de Arquivo)
	// e referir-se a uma confirmação de lançamento
	// P004
	// Tamanho: 15
	ValorReal uint64
}

type Moeda struct {
	// Códito adotado pela FEBRABAN para identificar a moeda utilizada
	// para expressar o valor do documento
	// *G040
	// Tamanho: 3
	Tipo TipoMoeda

	// Número de unidades do tipo de moeda identificada para cálculo do valor do documento
	// G041
	// Tamanho: 15
	Quantidade uint64
}

type TipoMoeda string

// Domínio dos tipo de moeda
const (
	TipoMoeda_BTN TipoMoeda = "BTN"
	TipoMoeda_BRL TipoMoeda = "BRL"
	TipoMoeda_USD TipoMoeda = "USD"
	TipoMoeda_PTE TipoMoeda = "PTE"
	TipoMoeda_FRF TipoMoeda = "FRF"
	TipoMoeda_CHF TipoMoeda = "CHF"
	TipoMoeda_JPY TipoMoeda = "JPY"
	TipoMoeda_IGP TipoMoeda = "IGP"
	TipoMoeda_IGM TipoMoeda = "IGM"
	TipoMoeda_GBP TipoMoeda = "GBP"
	TipoMoeda_ITL TipoMoeda = "ITL"
	TipoMoeda_DEM TipoMoeda = "DEM"
	TipoMoeda_TRD TipoMoeda = "TRD"
	TipoMoeda_UPC TipoMoeda = "UPC"
	TipoMoeda_UPF TipoMoeda = "UPF"
	TipoMoeda_UFR TipoMoeda = "UFR"
	TipoMoeda_XEU TipoMoeda = "XEU"
)

var TipoMoeda_Valor = map[string]string{
	"BTN": "Bônus do Tesouro Nacional + TR",
	"BRL": "Real",
	"USD": "Dólar Americano",
	"PTE": "Escudo Português",
	"FRF": "Franco Francês",
	"CHF": "Franco Suiço",
	"JPY": "Ien Japonês",
	"IGP": "Índice Geral de Preços",
	"IGM": "Índice Geral de Preços de Mercado",
	"GBP": "Libra Esterlina",
	"ITL": "Lira Italiana",
	"DEM": "Marco Alemão",
	"TRD": "Taxa Referencial Diária",
	"UPC": "Unidade Padrão de Capital",
	"UPF": "Unidade Padrão de Financiamento",
	"UFR": "Unidade Fiscal de Referência",
	"XEU": "Unidade Monetária Européia",
}

func CriarCredito(
	numeroEmpresa string,
	dataPagamento uint32,
	moeda Moeda,
	valorPagamento uint64,
	numeroBanco string,
) (Credito, error) {
	if len(numeroEmpresa) > 20 {
		return Credito{}, errors.New("NumeroEmpresa deve ter até 20 caracteres")
	} else if dataPagamento > 99999999 {
		return Credito{}, errors.New("DataPagamento deve ter até 8 dígitos")
	} else if valorPagamento > 999999999999999 {
		return Credito{}, errors.New("ValorPagamento deve ter até 15 dígitos")
	} else if len(numeroBanco) > 20 {
		return Credito{}, errors.New("NumeroBanco deve ter até 20 caracteres")
	}
	return Credito{numeroEmpresa, dataPagamento, moeda, valorPagamento, numeroBanco, 0, 0}, nil
}

func CriarMoeda(tipo TipoMoeda, quantidade uint64) (Moeda, error) {
	if _, ok := TipoMoeda_Valor[string(tipo)]; !ok {
		return Moeda{}, errors.New("TipoMoeda não encontrado")
	} else if quantidade > 999999999999999 {
		return Moeda{}, errors.New("Quantidade deve ter até 15 dígitos")
	}
	return Moeda{tipo, quantidade}, nil
}

func (c Credito) Processar() string {
	sNumeroEmpresa := fmt.Sprintf("%-20s", c.NumeroEmpresa)
	sDataPagamento := fmt.Sprintf("%08d", c.DataPagamento)
	sValorPagamento := fmt.Sprintf("%015d", c.ValorPagamento)
	sNumeroBanco := fmt.Sprintf("%-20s", c.NumeroBanco)
	sDataReal := fmt.Sprintf("%08d", 0)
	sValorReal := fmt.Sprintf("%015d", 0)
	return sNumeroEmpresa[:20] + sDataPagamento[:8] + c.Moeda.Processar() + sValorPagamento[:15] + sNumeroBanco[:20] + sDataReal[:8] + sValorReal[:15]
}

func (m Moeda) Processar() string {
	sTipo := fmt.Sprintf("%-3s", m.Tipo)
	sQuantidade := fmt.Sprintf("%015d", m.Quantidade)
	return sTipo[:3] + sQuantidade[:15]
}
