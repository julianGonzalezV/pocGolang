package main

import "log"

// Este es una practica comun, el de retornar channels
func doSomething() chan int {
	ch := make(chan int)
	go func() {
		for i := range ch {
			log.Println(i)
		}
	}()
	return ch
}

func main() {
	ch := doSomething()
	ch <- 1
	ch <- 4
	ch <- 8
}
