package main

import "fmt"

func main() {
	/// var a1 int8 = 128 compile error de over floww
	var a int8 = 125 // muy diferente a si inicia en un valor válido y luego lo incremente
	/// saca exception? R/ No, mire el 'for' que sigue
	var b uint8 = 253
	///al llegar al límite no se revienta peero saca valores negativos o ceros
	/// Tenga cuidado al asignar el tipo de valor adecuado
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
}
