package main

import (
	"fmt"
)

func main() {
	offset := 5
	fmt.Println(offset)
	offset = 10 // si usa :=  no compila porque es para creación mas no para acualización
	fmt.Println(offset)

	query, limit, offset := "bat", 10, 0
	fmt.Println(query, limit, offset)
	query, limit, offset = "ball", offset, 20 // note que también se usa '='
	fmt.Println(query, limit, offset)

	/* Shorthand Operators
		--: Reduce a number by 1
	++: Increase a number by 1
	+=: Add and assign
	-=: Subtract and assign*/
}
