package main

import (
	"fmt"
)

/// Variadic function permite cantidad variable de input parameters
// Note. if your function requires multiple parameters, the variadic parameter must be the last parameter in the function
func nums(i ...int) {
	// note que el tipo es un slice de int
	fmt.Printf("%T\n Len: %d\n  Cap: %d\n", i, len(i), cap(i))
	fmt.Println("Value => ", i)
}
func main() {
	fmt.Println("::::::::..Variadic function1:::::::")
	nums(99, 100)
	nums(200)
	nums()

	fmt.Println("::::::::..Variadic Passign slice to function:::::::")
	i := []int{5, 10, 15}
	// nums(i) error porque la funcion recibe 1  omas ints no []ints
	nums(i...) ///Ok esta es la forma de decirle que se el envía uno o mas enteros
}
