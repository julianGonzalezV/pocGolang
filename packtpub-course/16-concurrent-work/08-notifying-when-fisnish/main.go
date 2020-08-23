package main

import "log"

func readThem(in, out chan string) {
	for i := range in {
		log.Println(i)
	} // Sending notification saying that the process has finishes

	out <- "fin jeje" // Si quita esta señal o mensaje el main thread
	// ira a deadLock dado que se quedó esperando que algo llegara y no llegó

}

func main() {
	log.SetFlags(0) // flags en 0 para que solo muestre el valor sin horas
	in, out := make(chan string), make(chan string)
	go readThem(in, out)

	strs := []string{"a", "b", "c", "d", "e", "f"}
	for _, s := range strs {
		in <- s // sending each string to the 'in' channel
	}

	close(in) // close the receiver channel

	//The main function terminates only when it has been notified that all the incoming messages have been sent
	// mediante el envío de alguna señal cualquiera no tiene que se el string "Done" :)
	<-out // wait for done signal

	/*
		In this exercise, you've learned how you can make a Goroutine
		 notify another Goroutine that the work has finished by passing
		 a message through a channel without needing WaitGroup.

		Importante la anoación "without needing WaitGroup"
		que fué el código en 07-buffers, intente crear una nueva version del 07-buffer sin
		waitgroup
	*/
}
