package main

/// Ejemplo de como distribuir trabajo en varias go routines y luego
/// reunir/obtener todos los computos en un solo resultado en un solo goRoutine(el main)
import (
	"log"
	"time"
)

/// Sending number messages
func push(from, to int, out chan int) {
	for i := from; i <= to; i++ {
		out <- i
		time.Sleep(time.Microsecond)
	}
}

func main() {
	s1 := 0
	ch := make(chan int, 100)

	/// 4 go routines
	go push(1, 25, ch)
	go push(26, 50, ch)
	go push(51, 75, ch)
	go push(76, 100, ch)

	for c := 0; c < 100; c++ {
		i := <-ch
		log.Println(i)
		s1 += i
	}
	log.Println(s1)

}
