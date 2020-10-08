package cobranca

import (
	"testing"
)

func TestCodigoModalidae(t *testing.T) {
	var m CodigoModalidade
	got := m.String()
	expected := ""
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	got = Simples.String()
	expected = "1"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}

func TestSituacaoBoleto(t *testing.T) {
	var s SituacaoBoleto
	got := s.String()
	expected := ""
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	got = SituacaoBoletoNormal.String()
	expected = "01"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	got = SituacaoBoletoBaixado.String()
	expected = "07"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}

	got = SituacaoBoletoPagoParcialmente.String()
	expected = "18"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
