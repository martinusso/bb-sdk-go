package cobranca

import (
	"encoding/json"
	"time"
)

type RetornoBaixa struct {
	NumeroContratoCobranca string    `json:"numeroContratoCobranca"`
	DataBaixa              time.Time `json:"dataBaixa"`
	HorarioBaixa           time.Time `json:"horarioBaixa"`
}

func (b *RetornoBaixa) UnmarshalJSON(data []byte) error {
	type Alias RetornoBaixa
	aux := &struct {
		*Alias
		DataBaixa    string `json:"dataBaixa"`
		HorarioBaixa string `json:"horarioBaixa"`
	}{
		Alias: (*Alias)(b),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	b.DataBaixa, _ = time.Parse("02.01.2006", aux.DataBaixa)
	b.HorarioBaixa, _ = time.Parse("15:04:05", aux.HorarioBaixa)
	return nil
}

type ErrorBaixaBoleto struct {
	Errors []struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	} `json:"errors"`
}

func (e *ErrorBaixaBoleto) Error() string {
	if len(e.Errors) > 0 {
		return e.Errors[0].Message
	}
	return ""
}

func (e ErrorBaixaBoleto) Codigo() string {
	if len(e.Errors) > 0 {
		return e.Errors[0].Code
	}
	return ""
}
