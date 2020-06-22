package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var message = ""
	//message := "" // Short way
	count := 5
	if count > 5 {
		message = "Greater than 5"
	} else {
		message = "Not greater than 5"
	}
	fmt.Println(message)

	//Bad Count Bug
	count2 := 0
	if count2 < 5 {
		count2 = 10
		count2++
	}
	fmt.Println((-2) % 4)

	a := uint8(255)
	b := uint8(127)
	c := a &^ b
	fmt.Printf("Type: %T Value: %v\n", c, c)
	fmt.Println("Int", int(c))
	fmt.Println(a, "+", b, "=", c)
	fmt.Println(a, "+", b, "=", 255+127)

	type T2 struct {
		a int
		b int
	}
	t := T2{a: 123, b: 234}
	fmt.Println(unsafe.Sizeof(t))
	fmt.Println(unsafe.Sizeof(&t))
	const d0 int = 1
	// d0 = 2 canot assing to a

	const d1 = 5
	const d2 complex128 = d1
	fmt.Println(d1, d2)

	/*
			//11
			var string int = 1
			var d3 string = "a"
			fmt.Printf("%T %T", string, d3)

		  // 12
			var int string = "a"
			var d4 int = 1
			fmt.Printf("%T %T", int, d4)*/

	// LOS DEMAS DENTO DE CONST TOMAN EL VALOR DE W1 POR DEFECTO
	const (
		W1 = 1
		W2
		W3
	)
	fmt.Println("W1, W2, W3 = ", W1, W2, W3)

	const (
		W11 = iota + 1
		W21 = iota + 2
		W31 = iota + 3
	)
	fmt.Println(W11, W21, W31) // 1,3,5

	const (
		W12 = 1 << iota
		W22 = 1 << iota
		W32 = 1 << iota
	)
	fmt.Println(W12, W22, W32) // 1 2 4

	type ByteSize int64
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
	)
	fmt.Println(KB, MB, GB) // 1024 1048576 1073741824

	const (
		max1 = 10
	)
	const (
		a1 = (max1 - iota)
		b1
		c1
	)
	fmt.Println(a1, b1, c1) // 10 9 8

	const (
		x2 = string(iota + 'a')
		y2
		z2
	)
	fmt.Println(x2, y2, z2) // A B C
}
