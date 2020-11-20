package main

import (
	"fmt"
	"time"
)

// Value: Can be passed to functions, if the function changes its value, this change make effect only
// inside the function. because Go always make a copy before the function call
//this cause memory consumption---Values  reside on the STACK
//Pointer:  directions to a value you want, where you can only get the value and using it
// Go doesnt make copies with pointers  --Pointers put its values on the HEAP

// Garbage collector works on the HEAP(por eso si se manejan ok tienen mejor majeno de memoria)

// tIP
// To prevent runtime errors, you can compare a pointer to nil before
// trying to get its value. This looks like <pointer> != nil.
func main() {
	// formas de crear un puntero
	var count1 *int    // eg 1
	count2 := new(int) // eg 2

	countTemp := 5       // eg 3.1 --create a variable (continue with 3.2)
	count3 := &countTemp // eg 3.2 create a pointer Using &, create a pointer from the existing variable:

	t := &time.Time{} // eg5 from a struct ..in this case Time is a struct

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3 before: %#v\n", countTemp)
	fmt.Printf("count3 after: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)

	//Getting a Value from a Pointer
	// se logra usando * ante del nombre de la variable(conocido como dereference), ahora You don't always need to
	// dereference; for example, when a property or function is on a struct

	fmt.Println(" Getting a Value from a Pointer ")
	if count1 != nil {
		fmt.Printf("count1: %#v\n", *count1)
	}

	if count2 != nil {
		fmt.Printf("count2: %#v\n", *count2)
	}

	if count3 != nil {
		fmt.Printf("count3: %#v\n", *count3)
	}

	if t != nil {
		fmt.Printf("time : %#v\n", *t)
		fmt.Printf("time function : %#v\n", t.String()) // When calling a function of structs we don't need to dereference it
	}

}
