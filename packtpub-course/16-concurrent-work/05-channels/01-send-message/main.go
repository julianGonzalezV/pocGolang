package main

import "log"

// Channels es la vía por la cual diferentes Goroutines puedem compartir contenido

func greet(ch chan string) {
	ch <- "mensaje" // Escribiendo/enviando mensajes hacha el channel
}

func sendNumber(ch chan int) {
	ch <- 7 // Escribiendo/enviando mensajes hacha el channel
}

func main() {
	log.Println("Channel1 -->")
	ch := make(chan string)
	go greet(ch)
	log.Println(<-ch) // <-ch: leyendo/escuchando del channel nombrado ch

	log.Println("Channel2 -->")
	ch2 := make(chan int, 1) // 1 es el tamanio del buffer/capacidad, si lo quita entonces
	// saldrá error all goroutines are asleep - deadlock!..DUDA PORQUE EN EL channel de string
	// no es necesario el buffer?
	sendNumber(ch2)
	i := <-ch2
	log.Println(i)

}
