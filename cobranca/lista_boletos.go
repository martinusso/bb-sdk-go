package cobranca

import (
	"encoding/json"
	"net/url"
	"strconv"
	"time"
)

type BoletosListagem struct {
	IndicadorContinuidade string                `json:"indicadorContinuidade"`
	QuantidadeRegistros   int                   `json:"quantidadeRegistros"`
	ProximoIndice         int                   `json:"proximoIndice"`
	Boletos               []BoletosListagemItem `json:"boletos"`
}

type BoletosListagemItem struct {
	NumeroBoleto             string    `json:"numeroBoletoBB"`
	DataRegistro             time.Time `json:"dataRegistro"`
	DataVencimento           time.Time `json:"dataVencimento"`
	ValorOriginal            float64   `json:"valorOriginal"`
	CarteiraConvenio         int       `json:"carteiraConvenio"`
	VariacaoCarteiraConvenio int       `json:"variacaoCarteiraConvenio"`
	EstadoTituloCobranca     string    `json:"estadoTituloCobranca"`
	Contrato                 int64     `json:"contrato"`
	DataMovimento            time.Time `json:"dataMovimento"`
	ValorAtual               float64   `json:"valorAtual"`
	ValorPago                float64   `json:"valorPago"`
}

func (b *BoletosListagemItem) UnmarshalJSON(data []byte) error {
	type Alias BoletosListagemItem
	aux := &struct {
		*Alias
		DataRegistro   string `json:"dataRegistro"`
		DataVencimento string `json:"dataVencimento"`
		DataMovimento  string `json:"dataMovimento"`
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	b.DataRegistro, _ = time.Parse("02.01.2006", aux.DataVencimento)
	b.DataVencimento, _ = time.Parse("02.01.2006", aux.DataVencimento)
	b.DataMovimento, _ = time.Parse("02.01.2006", aux.DataMovimento)
	return nil
}

type ListaBoletosParams struct {
	IndicadorSituacao          IndicadorSituacao
	AgenciaBeneficiario        string
	ContaBeneficiario          string
	CarteiraConvenio           int
	VariacaoCarteiraConvenio   int
	ModalidadeCobranca         CodigoModalidade
	CnpjPagador                string
	DigitoCnpjPagador          string
	CpfPagador                 string
	DigitoCpfPagador           string
	DataInicioVencimento       time.Time
	DataFimVencimento          time.Time
	DataInicioRegistro         time.Time
	DataFimRegistro            time.Time
	DataInicioMovimento        time.Time
	DataFimMovimento           time.Time
	CodigoEstadoTituloCobranca SituacaoBoleto
	BoletoVencido              BoletoVencido
	Indice                     int
}

func (p ListaBoletosParams) Values() map[string][]string {
	params := url.Values{}

	params.Set("indicadorSituacao", string(p.IndicadorSituacao))
	params.Set("agenciaBeneficiario", p.AgenciaBeneficiario)
	params.Set("contaBeneficiario", p.ContaBeneficiario)

	if p.CarteiraConvenio > 0 {
		params.Set("carteiraConvenio", strconv.Itoa(p.CarteiraConvenio))
	}
	if p.VariacaoCarteiraConvenio > 0 {
		params.Set("variacaoCarteiraConvenio", strconv.Itoa(p.VariacaoCarteiraConvenio))
	}
	if p.ModalidadeCobranca > 0 {
		params.Set("modalidadeCobranca", p.ModalidadeCobranca.String())
	}
	if p.CnpjPagador != "" {
		params.Set("cnpjPagador", p.CnpjPagador)
	}
	if p.DigitoCnpjPagador != "" {
		params.Set("digitoCnpjPagador", p.DigitoCnpjPagador)
	}
	if p.CpfPagador != "" {
		params.Set("cpfPagador", p.CpfPagador)
	}
	if p.DigitoCpfPagador != "" {
		params.Set("digitoCpfPagador", p.DigitoCpfPagador)
	}
	if !p.DataInicioVencimento.IsZero() {
		params.Set("dataInicioVencimento", p.DataInicioVencimento.Format("02.01.2006"))
	}
	if !p.DataFimVencimento.IsZero() {
		params.Set("dataFimVencimento", p.DataFimVencimento.Format("02.01.2006"))
	}
	if !p.DataInicioRegistro.IsZero() {
		params.Set("dataFimRegistro", p.DataFimRegistro.Format("02.01.2006"))
	}
	if !p.DataFimRegistro.IsZero() {
		params.Set("dataFimRegistro", p.DataFimRegistro.Format("02.01.2006"))
	}
	if !p.DataInicioMovimento.IsZero() {
		params.Set("dataFimMovimento", p.DataFimMovimento.Format("02.01.2006"))
	}
	if !p.DataFimMovimento.IsZero() {
		params.Set("dataFimMovimento", p.DataFimMovimento.Format("02.01.2006"))
	}
	if p.CodigoEstadoTituloCobranca > 0 {
		params.Set("codigoEstadoTituloCobranca", p.CodigoEstadoTituloCobranca.String())
	}
	if p.BoletoVencido != "" {
		params.Set("boletoVencido", string(p.BoletoVencido))
	}
	if p.Indice > 0 {
		params.Set("indice", strconv.Itoa(p.Indice))
	}
	return params
}
