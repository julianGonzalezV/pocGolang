package main

import (
	"errors"
	"fmt"
)

/// Go no usa excepciones y en su lugar usa panic que es una funcion que causa que el
/// programa se estalle ;)

/*
Cuando un panic ocurre, entonces se ejecutan estos pasos:

The execution is stopped
Any deferred functions in the panicking function will be called
Any deferred functions in the stack of the panicking function will be called
It will continue all the way up the stack until it reaches main()
Statements after the panicking function will not execute
The program then crashes
*/

// Algunas ejeciones que causan un panic (condiciones anormales que se agregarian al programa)
// El famoso index out of range

func crash1(msg string) {
	if msg == "good-bye" {
		panic(errors.New("something went wrong"))
	}
}

func main() {
	fmt.Println("Panic: When occurs then code below panic function won execute")
	msg := "good-bye"
	crash1(msg)
	fmt.Println("This line will not get printed")
}
