package main

import "log"

/// push: have in chan for incomes request and out chan for sending messages
func push(from, to int, in chan bool, out chan int) {
	for i := from; i <= to; i++ {
		boolMsg := <-in // waiting for a request
		if boolMsg {
			out <- i // when request(previous line) is received  as true value then 'i' is sent as message
		} else {
			out <- 0
		}

	}
}

func main() {
	// Similar to summ created in folder 03-sum-numbers but this one creates a new channel 'in'
	s1 := 0
	out := make(chan int, 100)
	in := make(chan bool, 100)
	go push(1, 25, in, out)
	go push(26, 50, in, out)
	go push(51, 75, in, out)
	go push(76, 100, in, out)

	for c := 0; c < 100; c++ {
		in <- true
		i := <-out
		log.Println(i)
		s1 += i
	}
	log.Println(s1)

}
