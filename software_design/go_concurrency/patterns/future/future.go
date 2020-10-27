package future

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

type MaybeString struct{}

func (s *MaybeString) Success(f SuccessFunc) *MaybeString {
	return nil
}

func (s *MaybeString) Fail(f FailFunc) *MaybeString {
	return nil
}

func (s *MaybeString) Execute(f ExecuteStrFunc) {

}
