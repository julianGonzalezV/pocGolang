package main

import (
	"fmt"
	"time"
)

func next(v *int) {
	c := *v
	*v = c + 1
}

func main() {

	a := 0
	next(&a)
	next(&a)
	next(&a)
	fmt.Println("a => ", a)

	b := 0
	go next(&b)
	go next(&b)
	go next(&b)
	time.Sleep(1000)
	fmt.Println("b => ", b) //b holds 3, or 2, or 1. Why would this happen?
	// El comportamiento anterior se dá porque cada go routine 'go next(&b)'
	// lo que hace es incrementear 'v' sin tener en cuenta(ni sabe) que valor va
	// esto es RACE CONDITION

}
