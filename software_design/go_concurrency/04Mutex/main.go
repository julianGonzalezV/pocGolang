package main

func main() {
	/// MUTEX
	CounterConcurrent()

	// Ejecutando el comando $run -race 04mutexes.go main.go
	// podrá ver si hay problemas de race condition

	// Si Desea verificar el funcionamiento del comand vaya a counterConcurrent()
	// y quite el lock()  y unlock() si solo quite el primero saldrá error "fatal error: sync: unlock of unlocked mutex"
	// si solo comenta el  unlock() se bloquea el applicativo, al parecer no encuentra que desbloquear
	// comentando ambos por fin verá el WARNING: DATA RACE, note que igual el resultado se devuelve correctamente

}
