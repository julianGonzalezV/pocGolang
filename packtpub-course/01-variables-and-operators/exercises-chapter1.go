package main

import (
	"fmt"
)

/// swapEg1 does not require third aux variable
func swapEg1() {
	fmt.Print("Enter first number : ")
	var first int
	fmt.Scanln(&first)
	fmt.Print("Enter second number : ")
	var second int
	fmt.Scanln(&second)
	first = first - second
	fmt.Println("op1 :", first)
	second = first + second
	fmt.Println("op2 :", second)

	first = second - first
	fmt.Println("op3 :", first)
	fmt.Println("First number :", first)
	fmt.Println("Second number :", second)
}

func swapEg2(a *int, b *int) {
	// swap the values here
	// R/
	var aux = *a // dereference
	fmt.Println("dereference 1 :", *a)
	fmt.Println("Pointer :", a)
	fmt.Println("dereference :", aux)
	fmt.Println("Pointer  :", &aux)

	*a = *b // normal
	fmt.Println("normal :", a)
	*b = *&aux // &aux creates a pointer based on aux variable(l28). After that, * get the value/dereference
	fmt.Println("b1 :", b)
	fmt.Println("b2 :", *b)
}

// swapEg3 does not saw anything because they are local variables into this functions
// and as a resutl it does not affect external world
func swapEg3(a int, b int) {
	var aux = a
	a = b
	b = aux
}

func main() {
	swapEg1()

	// call swap here
	// 1: Call the swap function, ensuring you are passing a pointer.
	a, b := 5, 10
	fmt.Println("INPUT a , b = ", a, ",", b)
	swapEg2(&a, &b)
	fmt.Println("OUTPUT a , b = ", a, ",", b)
	fmt.Println(a == 10, b == 5)

	x, y := 5, 10
	fmt.Println("INPUT x , y = ", x, ",", y)
	swapEg3(x, y)
	fmt.Println("OUTPUT x , y = ", x, ",", y)

	z := 10
	z = 15
	fmt.Println("OUTPUT z = ", z)

}
