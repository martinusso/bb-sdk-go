package cobranca

import (
	"testing"
)

func TestNossoNumero(t *testing.T) {
	got := NossoNumero("9876543", "123")
	expected := "00098765430000000123"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
