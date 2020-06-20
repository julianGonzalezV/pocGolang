package main

import (
	"fmt"
)

func add5Value(count int) {
	count += 5
}

func add5Point(count *int) {
	*count += 5
}

func main() {
	var count int
	add5Value(count)
	fmt.Println("add5Value   :", count) // note como no lo afecta, forma recomendada de usar pero
	// aveces afecta el diseño por un lado o el performance por consumo de memoria ya que siempre se
	// hace una copia en memoria ..aunque diría que las memorias hoy en día no debería ser un problema

	add5Point(&count)                   // &count crea un puntero a partir de la variable count ( linea 16 sin *)
	fmt.Println("add5Point   :", count) // note como lo modificó, como paso por referencia en JAVA ..se hace peligroso
	// es verdad que ahorra memoria pero OJO con su uso
}
