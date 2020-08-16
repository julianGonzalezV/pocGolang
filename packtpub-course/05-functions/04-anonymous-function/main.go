package main

import "fmt"

func main() {
	func() {
		fmt.Println("Anonymous #1")
	}() //() se conoce colo los parentesis ejecutores

	fmt.Println("Anonymous #2 - parameters")
	message := "Paol"
	func(str string) {
		fmt.Println(str)
	}(message)

	fmt.Println("Anonymous #3::::::::::: ")
	f1 := func() {
		fmt.Println("Assing to variable")
	}
	// Calling
	f1()
	fmt.Println("Anonymous #4::::::::::: ")
	f2 := func(inSt string) {
		fmt.Println(inSt)
	}
	// Calling
	f2("Soy un parametro!!")

	fmt.Println("Anonymous #4 Square::::::::::: ")
	x := func(i int) int {
		return i * i
	}
	fmt.Printf("Square resutl %d\n", x(9))
}
