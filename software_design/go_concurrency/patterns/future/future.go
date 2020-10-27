package future

import "fmt"

/// Also called PROMISES
///
/// Useful for  asynchronous programing; Write algorithms that will be executed
/// or not in the future by the same goRoutine or another one

/// fire and forget: Estrategia que lo que hace es ejecutar el computo en un futuro
/// Dentro de un nuevo goRoutine(diferente al principal) y NO espera por alguna respuesta
/// ni por el status de progreso, lo que se define es que hacer en caso de EXITO y en caso de Error
/// y chao

///function types
type SuccessFunc func(string)
type FailFunc func(error)
type ExecuteStrFunc func() (string, error)

type MaybeString struct {
	successFunc SuccessFunc
	failFunc    FailFunc
}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	s.successFunc = f // storing success func f into s variable
	return s          /// al retornar el mismo *MaybeString permite el encadenamiento de computo
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	s.failFunc = f
	return s
}

/// Execute takes a ExecuteStrFunc and execute it asynchronously
func (s *MaybeString) Execute(f ExecuteStrFunc) {
	/// Delegating to a goRoutine the result when SUCESS or FAILURE
	go func(inputMs *MaybeString) {
		str, err := f()
		if err != nil {
			s.failFunc(err)
		}
		s.successFunc(str)
	}(s)
}

func setContext(msg string) ExecuteStrFunc {
	msg = fmt.Sprintf("%s Closure!\n", msg)
	/// todo lo que está acá afuera no se correrá concurrentemente
	/// y será ejecutado asincrónicamente antes de llamar la sgte funcion anónima
	/// POR ESO trate de llevar la logica dentro de la función anónima
	return func() (string, error) {
		return msg, nil
	}
}
