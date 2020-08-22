package main

import (
	"fmt"
	"log"
)

func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Recibiendo mensaje  %s", msg)
	ch <- "Enviando mensaje"
}

func main() {
	ch := make(chan string)
	go greet(ch)       // Second goRoutine is started
	ch <- "como fuaka" // Sending message from the main Routine(thread in other languages)
	/// The second routine is waiting/listening
	log.Println(<-ch)
	log.Println(<-ch)
	// Debe usar Println tantos mensajes envié ..por lo regular lo que hacer es un loop para recibirlos todos (ver ejemplo en sum-numbers)
	// Note que se usa el mismo canal ch

}
