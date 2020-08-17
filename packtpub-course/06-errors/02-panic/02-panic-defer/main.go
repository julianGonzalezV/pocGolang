package main

// Cuando se tienen defer functions y un panic se ejecuta entonces primero se ejecutan los
// defers y por último el panic

// Existe otra funcion---> os.Exit() esta si para todo el sistema inmediatamente sin
// permitir la ejecución del los defers

/// recuerde que los defer se almacenan en una especie de estructura stack
/// on ejecución LiFo cuando llega el momento de ejecutarlas todas
import (
	"errors"
	"fmt"
)

func main() {
	test()
	fmt.Println("This line will not get printed")
}
func test() {
	n := func() {
		fmt.Println("Defer in test")
	}
	defer n()
	msg := "good-bye"
	message(msg)
}
func message(msg string) {
	f := func() {
		fmt.Println("Defer in message func")
	}
	defer f()
	if msg == "good-bye" {
		panic(errors.New("something went wrong"))
	}
}
