package main

/// The important thing to remember is that Go is not a parallel language
/// It implements coroutines   eg: Hammer nails example
import (
	"fmt"
	"time"
)

func printSomething(sm string) {
	fmt.Println(sm)
}

func sum(from, to int) int {
	res := 0
	for i := from; i <= to; i++ {
		res += i
	}
	return res
}

// Main routine
func main() {
	fmt.Println("Start")
	go printSomething("Other process")
	fmt.Println("End")

	fmt.Println("::::::.Sum sin gorutine")
	s1 := sum(1, 10)
	fmt.Println("Sum sin gorutine => ", s1)

	fmt.Println("::::::.Sum CON gorutine")
	var s2 int
	go func() {
		s2 = sum(1, 10)
	}()

	time.Sleep(time.Second) /// sin el timer la respueta sería 0
	// El timer solo es por temas de pruebas :)
	fmt.Println("Sum Con gorutine => ", s2)
}
