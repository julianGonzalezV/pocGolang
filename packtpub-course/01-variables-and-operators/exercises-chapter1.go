package main

import (
	"fmt"
)

func swap(a *int, b *int) {
	// swap the values here
	// R/
	var aux = *a // dereference
	*a = *b      // normal
	*b = *&aux   // &aux creates a pointer based on aux variable(L10). After that, * get the value/dereference
}

func main() {
	a, b := 5, 10
	// call swap here
	// 1: Call the swap function, ensuring you are passing a pointer.
	// R/
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)

}
