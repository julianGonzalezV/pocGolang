package main

import "fmt"

/// Regular functions cannot reference variables outside of themselves
/// but anonymous can
/// Note el tipo de retorno ---> una funcion anonima!
func incrementorF() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func decrement(i int) func() int {
	return func() int {
		i -= 1
		return i
	}

}

func main() {
	i := 0
	incrementor := func() int {
		i += 1
		return i
	}
	fmt.Println(incrementor())
	fmt.Println(incrementor())
	i += 10
	fmt.Println(incrementor())

	fmt.Println("Increment 1:::::::")
	increment := incrementorF()
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println("Increment 2:::::::")
	/// Note que lo interesante es que i dentro de incrementorF no se afecta por la
	/// Creación de increment2, es thread safe
	increment2 := incrementorF()
	fmt.Println(increment2())
	fmt.Println(increment2())

	fmt.Println("Decrement 1:::::::")
	int1 := 4
	decrement1 := decrement(int1)
	fmt.Println(decrement1())
	fmt.Println(decrement1())

	fmt.Println("Decrement 2:::::::")
	/// Mas interesante aún!!, note como a pesar en de enviarle el mismo
	/// int1 que se crea e inicializa en el main , no afecta!!
	decrement2 := decrement(int1)
	fmt.Println(decrement2())
	fmt.Println(decrement2())
}
