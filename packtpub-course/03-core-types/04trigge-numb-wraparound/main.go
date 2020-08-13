package main

import (
	"fmt"
	"math"
	"math/big"
)

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

	intA := math.MaxInt64
	intA = intA + 1

	/// BIG NUMBERS
	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))
	fmt.Println("MaxInt64: ", math.MaxInt64) // note como retorna el numero negativo
	fmt.Println("Int   :", intA)
	fmt.Println("Big Int : ", bigA.String())

	//The byte type in Go is just an alias for uint8

	/// TEXT Raw(usin ``) and Interpreted(using "") Strings literals
	/// Note como con ` se imprime tal cual como se esqcribe, espacios, saltos de líena etc
	comment1 := `This is the BEST 
thing ever!`
	comment2 := `This is the BEST\nthing ever!`
	comment3 := "This is the BEST\nthing ever!"
	fmt.Print(comment1, "\n\n")
	fmt.Print(comment2, "\n\n")
	fmt.Print(comment3, "\n")

	// Note como al usar Raw String hace que se lea mejor que con la version Interpreted
	comment4 := `In "Windows" the user directory is "C:\Users\"`
	comment5 := "In \"Windows\" the user directory is \"C:\\Users\\\""
	fmt.Println(comment4)
	fmt.Println(comment5)

	/// Rune: type to storage a single UTF() multi-byte character

	username := "Sir_King_Über"
	fmt.Println("Print every positions on username var")
	for i := 0; i < len(username); i++ {
		fmt.Print(username[i], " ")
	}
	fmt.Println("Converting to string and Printing every positions on username var")
	/// Note como acá omite el Ü porque este fue codificado usando algo maś que Byte
	for i := 0; i < len(username); i++ {
		fmt.Print(string(username[i]), " ")
	}

	fmt.Println(" Usinf RUNE : Converting to string and Printing every positions on username var")
	/// El tipo Rune evita el error anterior con Ü, ya ques Ü es codiicado con multi-byte entonces
	/// y Rune permite este almacenamieto de single o mutubyte codificacoin
	/// lo que usamos por lo regualr con single
	runes := []rune(username)
	for i := 0; i < len(runes); i++ {
		fmt.Print(string(runes[i]), " ")
	}
}
