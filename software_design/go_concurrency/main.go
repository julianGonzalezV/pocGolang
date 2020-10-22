package main

import (
	"fmt"
)

func main() {
	GoRoutinesWithTime()
	/// Go routines using waitGroup instead of Time
	WaitG1()
	WaitG2()
	WaitG3()

	///Callbacks; anonymous function that will be executed
	/// within the context of a different function
	anonymousSynch := func(v string) {
		fmt.Printf("Result: %s\n", v)
	}
	// synchronous callback
	toUpperSync("salir en mayúscula!", anonymousSynch)

	// Asynchronous callback
	// usandolo conscientemente es potente pero OJO con el CALLBACK hell
	toUpperASyncCaller("salir en mayúscula Asynch!")

}
