/*
Te concurrecy model used by golang is Communicating Sequential Processes (CSP),
by contrast Actor model/system is used in languages like scala, erlang

The actor system requires actors to know the processIds for sending messages eachother
CSP no!, cause it uses channels, the way to communicate between processes
because channels are completely anonymous

::::::::Goroutines:::::::::::::::.
processes that run applications in a computer concurrently,They execute some logic and die (or keep looping if necessary).
Similar to Actors in the "Actor System"

THEY ARE NOT THREADS!! due to we can launch thousands or millions of concurrent goRoutines


*/

package main

import (
	"fmt"
	"time"
)

func messageProcess1() {
	fmt.Println("Executing go routine #1")

}

/// Note que al final se le coloca un Time
func GoRoutinesWithTime() {
	/// Go routines (they are processes but no threads!)
	// Creation
	/// 1) Normal function
	//time.Sleep(3 * time.Second) /// el timer no va antes! ya que lo haría
	// sobre la ejecución del routine actual(El main) y al despertarlo termina tan rápido
	// que no se alcanza a ejecutar la rutina messageProcess1
	go messageProcess1()

	/// 2) Anonymous function
	go func() {
		fmt.Println("Executing go routine #2")
	}()

	/// 3) Anonymous function with parameters
	go func(msg string) {
		println(msg)
	}("Executing go routine #3") // acá se le envía el parametro!

	/// 4) assigning Anonymous function to variables
	fAssign1 := func() {
		println("Executing go routine #4")
	}
	go fAssign1()

	fAssign2 := func(msg string) {
		println(msg)
	}
	go fAssign2("Executing go routine #5")

	time.Sleep(2 * time.Second) // este time es del hilo main

}
