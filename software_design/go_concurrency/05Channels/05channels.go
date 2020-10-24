package main

import (
	"fmt"
	"sync"
	"time"
)

// Channels(otra funcionalidad para escribir aplicaciones concurrentes)
// - Medio de comunicación entre procesos
// - ofrece una forma más natural y trasparente de implementar
// concurrancia en golang
// -En la mayoría de veces terminamos con un mejor diseño si lo comparamos
// con Mutex(mucho código y error prone)
// Una o mas goRoutines pueden emitir y uno o más Groutines pueden recivor
// como los programas de Tv :) (por cierto se emiten en dftes canales jeje)

func channel1() {
	channel := make(chan string) // make(type) data type the channel will transport
	go func() {
		channel <- "ok aquí vamos"
	}()
	message := <-channel // arrow before the channel name indicates that data will be extracted
	fmt.Println(message)
}

// What happens without receiver??
func channelNoRecipient() {
	channel := make(chan string) // make(type) data type the channel will transport
	go func() {
		channel <- "Data 2"
	}()
	//message := <-channel // arrow before the channel name indicates that data will be extracted
	//f
	//fmt.Println(message)
}

// What happens without receiver??
func channelNoRecipientV2() {
	var wait sync.WaitGroup
	channel := make(chan string) // make(type) data type the channel will transport
	wait.Add(1)
	go func() {
		channel <- "Data 3"
		println("Finishing channelNoRecipientV2")
		wait.Done()
	}()
	// Note que no se reciben mensajes hasta que pasen los segundos de espera
	//
	time.Sleep(2 * time.Second)
	message := <-channel // arrow before the channel name indicates that data will be extracted
	fmt.Println("channelNoRecipientV2 received", message)
	wait.Wait()
}

// Con buffer note que a comparación de channelNoRecipientV2() no se bloquea el sender
// el sender es el goRoutine (anonymous function), vea como println("Finishing bufferedChannel")
// se imprime siempre antes de que el recipient reciba el dato
func bufferedChannel() {
	var wait sync.WaitGroup
	channel := make(chan string, 1) //1
	wait.Add(1)
	go func() {
		channel <- "Data 4"
		println("Finishing bufferedChannel")
		wait.Done()
	}()
	time.Sleep(2 * time.Second)
	message := <-channel // arrow before the channel name indicates that data will be extracted
	fmt.Println("received => ", message)
	wait.Wait()
}

// a comparación de bufferedChannel() se envía 2 datos al channel pero como el buffer solo es 1 entonces:
// - Se envía Data5 y Data6
// - El gouRoutine se bloquea porque el buffer solo soporta 1 dato
// - Se desblquea (el sleep se termina)
// - se imprime Data 5 , Data6 y fisnis
func bufferedChannel2() {
	var wait sync.WaitGroup
	channel := make(chan string, 1)
	wait.Add(1)
	go func() {
		channel <- "Data 5"
		channel <- "Data 6"
		println("Finishing bufferedChannel, Data 6 debe llegar después de finalizar ")
		wait.Done()
	}()
	time.Sleep(2 * time.Second)
	message := <-channel // arrow before the channel name indicates that data will be extracted
	fmt.Println("received => ", message)
	wait.Wait()
}

/// For managing many incoming channels in the same goRoutine
func selectingChannels(ch1, ch2 <-chan string, ch3 chan<- bool) {
	for {
		select {
		case msg := <-ch1:
			println(msg)
		case msg := <-ch2:
			println(msg)
		case <-time.After(time.Second * 2):
			println("Nothing received in 2 seconds. Exiting")
			ch3 <- true
			break
		}
	}
}

func sendString(ch chan<- string, s string) {
	ch <- s
}

func selectAux() {
	helloCh := make(chan string, 1)
	goodbyeCh := make(chan string, 1)
	quitCh := make(chan bool)
	go selectingChannels(helloCh, goodbyeCh, quitCh)
	go sendString(helloCh, "hello!")
	time.Sleep(time.Second)
	go sendString(goodbyeCh, "goodbye!")
	<-quitCh
}

func rangingChannels() {
	fmt.Println("rangingChannels => ")
	ch := make(chan int)
	go func() {
		ch <- 1
		time.Sleep(time.Second)

		ch <- 2
		close(ch)

	}()
	for v := range ch { // the range iterates all the time until the goroutine close the channel
		println(v)
	}

}
