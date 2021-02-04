package errorinterface

import (
	"errors"
	"testing"
)

func Test1(t *testing.T) {
	de := New("Fallo en fecha")
	dePointer := NewPointer("DePointer Fallo en fecha")
	golangError := errors.New(" native Go error Fallo en fecha")
	/// Note que -de- es de tipo error(propio de golang) y golangError es un error (errors.New)
	/// y al llamar PrintError cumple con la firma
	de.PrintError(de) // si -de- no implementa Error() entonces PrintError falla alno recibor un tipo error
	de.PrintError(golangError)
	// dePointer es un puntero retornado en NewPointer method
	// por què funciona al pasar a PrintError??..R/ Porque INTERFACES are ALWAYS passed as pointers without having to stablish directly
	// goolan infers it has to be a pointer
	de.PrintError(dePointer)

}
