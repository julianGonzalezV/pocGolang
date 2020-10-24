package main

import (
	"fmt"
	"time"
)

func directional1() {
	println("directional1() => :")
	channel := make(chan string, 1)
	go func(ch chan<- string) { // arrow after word chan..chan<- idicates/restrict
		//the channel passed to this function to only be used for passing messages
		ch <- "string string string "
		println("Finishing directional1 goroutine")
	}(channel) // passing channel as parameter
	time.Sleep(time.Second)
	message := <-channel
	fmt.Println(message)
}

func receiverChannel(ch <-chan string) {
	msg := <-ch
	println(msg)
	// ch <- "hello"no va a funcionr porque el parámetro dice que solo se reciben channel que escuchen
}
