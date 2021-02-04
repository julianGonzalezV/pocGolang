package math_ops

import "testing"

func TestFactorial(t *testing.T) {
	factorR, expected := factorial(4), 24
	if factorR != expected {
		t.Errorf("Result %d but was expected %d \n", factorR, expected)
	}
}
