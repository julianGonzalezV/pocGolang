package main

import (
	"fmt"
	"time"
)

func next(v *int) {
	fmt.Println("1-v => ", v)
	fmt.Println("1- * v => ", *v)
	fmt.Println("1- &v => ", &v)
	c := *v
	fmt.Println("c => ", c)
	*v = c + 1
	fmt.Println("2-v => ", v)
	fmt.Println("2- * v => ", *v)
	fmt.Println("2- &v => ", &v)
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
	time.Sleep(time.Second)
	fmt.Println("b => ", b) //b holds 3, or 2, or 1. Why would this happen?
	// El comportamiento anterior se dá porque cada go routine 'go next(&b)'
	// lo que hace es incrementear 'v' sin tener en cuenta(ni sabe) que valor va
	// esto es RACE CONDITION

}
