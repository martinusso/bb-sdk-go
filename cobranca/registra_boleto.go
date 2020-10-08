package cobranca

import (
	"encoding/json"
	"errors"
	"strings"
	"time"
)

type Boleto struct {
	NumeroConvenio                       string           `json:"numeroConvenio"`
	NumeroCarteira                       int              `json:"numeroCarteira"`
	NumeroVariacaoCarteira               int              `json:"numeroVariacaoCarteira"`
	CodigoModalidade                     CodigoModalidade `json:"codigoModalidade"`
	DataEmissao                          time.Time        `json:"dataEmissao"`
	DataVencimento                       time.Time        `json:"dataVencimento"`
	ValorOriginal                        float64          `json:"valorOriginal"`
	ValorAbatimento                      float64          `json:"valorAbatimento"`
	QuantidadeDiasProtesto               int              `json:"quantidadeDiasProtesto"`
	IndicadorNumeroDiasLimiteRecebimento string           `json:"indicadorNumeroDiasLimiteRecebimento"`
	NumeroDiasLimiteRecebimento          int              `json:"numeroDiasLimiteRecebimento"`
	CodigoAceite                         CodigoAceite     `json:"codigoAceite"`
	CodigoTipoTitulo                     string           `json:"codigoTipoTitulo"`
	DescricaoTipoTitulo                  string           `json:"descricaoTipoTitulo"`
	IndicadorPermissaoRecebimentoParcial string           `json:"indicadorPermissaoRecebimentoParcial"`
	NumeroTituloBeneficiario             string           `json:"numeroTituloBeneficiario"`
	TextoCampoUtilizacaoBeneficiario     string           `json:"textoCampoUtilizacaoBeneficiario"`
	NumeroTituloCliente                  string           `json:"numeroTituloCliente"`
	TextoMensagemBloquetoOcorrencia      string           `json:"textoMensagemBloquetoOcorrencia"`
	Desconto                             Desconto         `json:"desconto"`
	JurosMora                            JurosMora        `json:"jurosMora"`
	Multa                                Multa            `json:"multa"`
	Pagador                              Pagador          `json:"pagador"`
	Avalista                             Avalista         `json:"avalista"`
	Email                                string           `json:"email"`
	QuantidadeDiasNegativacao            int              `json:"quantidadeDiasNegativacao"`
}

func (b Boleto) Validate() error {
	if len(strings.TrimSpace(b.NumeroConvenio)) != 7 {
		return errors.New("Número do convênio de cobrança inválido.")
	}
	if b.DataEmissao.Truncate(24 * time.Hour).Before(time.Now().Truncate(24 * time.Hour)) {
		//Qualquer data, a partir da data atual no formato “dd.mm.aaaa”, entre aspas. CAMPO OBRIGATÓRIO.
		return errors.New("Data de emissão deve ser a partir da data atual.")
	}
	if b.DataVencimento.Truncate(24 * time.Hour).Before(b.DataEmissao.Truncate(24 * time.Hour)) {
		//Qualquer data >= dataEmissao, no formato “dd.mm.aaaa”, entre aspas. CAMPO OBRIGATÓRIO
		return errors.New("Data de vencimento deve ser maior, ou igual, a data de emissão.")
	}
	return nil
}

func (b Boleto) MarshalJSON() ([]byte, error) {
	type Alias Boleto
	return json.Marshal(&struct {
		Alias
		DataEmissao    string `json:"dataEmissao"`
		DataVencimento string `json:"dataVencimento"`
	}{
		Alias:          (Alias)(b),
		DataEmissao:    b.DataEmissao.Format("02.01.2006"),
		DataVencimento: b.DataVencimento.Format("02.01.2006"),
	})
}

type RegistroBoleto struct {
	Numero                 string `json:"numero"`
	NumeroCarteira         int    `json:"numeroCarteira"`
	NumeroVariacaoCarteira int    `json:"numeroVariacaoCarteira"`
	CodigoCliente          int64  `json:"codigoCliente"`
	LinhaDigitavel         string `json:"linhaDigitavel"`
	CodigoBarras           string `json:"codigoBarraNumerico"`
	NumeroContratoCobranca int64  `json:"numeroContratoCobranca"`
	Beneficiario           struct {
		Agencia              int64  `json:"agencia"`
		ContaCorrente        int64  `json:"contaCorrente"`
		TipoEndereco         int    `json:"tipoEndereco"`
		Logradouro           string `json:"logradouro"`
		Bairro               string `json:"bairro"`
		Cidade               string `json:"cidade"`
		CodigoCidade         int64  `json:"codigoCidade"`
		UF                   string `json:"uf"`
		CEP                  int64  `json:"cep"`
		IndicadorComprocavao string `json:"indicadorComprovacao"`
	} `json:"beneficiario"`
	QuantidadeOcorrenciasNegativacao string `json:"quantidadeOcorrenciasNegativacao"`
	//"listaOcorrenciasNegativacao": []
}

type Desconto struct {
	Tipo TipoDesconto `json:"tipo"`
}

type JurosMora struct {
	Tipo TipoJurosMora `json:"jurosMora"`
}

type Multa struct {
	Tipo TipoMulta `json:"multa"`
}

type Avalista struct {
	Nome           string       `json:"nomeRegistro"`
	NumeroRegistro string       `json:"numeroRegistro"`
	TipoRegistro   TipoRegistro `json:"tipoRegistro"`
}

func (a Avalista) MarshalJSON() ([]byte, error) {
	type Alias Avalista
	return json.Marshal(&struct {
		Alias
		Nome string `json:"nomeRegistro"`
	}{
		Alias: (Alias)(a),
		Nome:  substring(a.Nome, 30),
	})
}

type Pagador struct {
	Nome           string       `json:"nome"`
	NumeroRegistro string       `json:"numeroRegistro"`
	TipoRegistro   TipoRegistro `json:"tipoRegistro"`
	Telefone       string       `json:"telefone"`
	Endereco       string       `json:"endereco"`
	Bairro         string       `json:"bairro"`
	CEP            string       `json:"cep"`
	Cidade         string       `json:"cidade"`
	UF             string       `json:"uf"`
}

func (p Pagador) MarshalJSON() ([]byte, error) {
	type Alias Pagador
	return json.Marshal(&struct {
		Alias
		Nome     string `json:"nome"`
		Endereco string `json:"endereco"`
		Bairro   string `json:"bairro"`
		Cidade   string `json:"cidade"`
	}{
		Alias:    (Alias)(p),
		Nome:     substring(p.Nome, 30),
		Endereco: substring(p.Endereco, 30),
		Bairro:   substring(p.Bairro, 30),
		Cidade:   substring(p.Cidade, 30),
	})
}

func substring(s string, l int) string {
	if len(s) > l {
		return s[0:l]
	}
	return s
}
