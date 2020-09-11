package operacionMat

import "testing"

/// TestMultiplyV1 inicialmente el test falla porque apenas tendremos la deinicion de la
/// funcion (att TDD :) )
func TestMultiplyV1(t *testing.T) {
	a, b, expected := 5, 6, 30
	if multiplyV1(a, b) != expected {
		t.Errorf("Our multiply function doens't work, %d*%d must be %d\n", a,
			b, expected)
	}
}

/// TestMultiply cuando pasa entonces no saldrá en los errores al
/// ejecutar go tes, ya que este comando solo devuelve los test que fallaron
func TestMultiply(t *testing.T) {
	a, b, expected := 5, 6, 30
	if multiply(a, b) != expected {
		t.Errorf("Our multiply function doens't work, %d*%d must be %d\n", a,
			b, expected)
	}
}
