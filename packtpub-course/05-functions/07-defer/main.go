package main

///The deferred functions are commonly used for performing "clean-up" activities
import "fmt"

/// Defer
func done() {
	fmt.Println("Now I am done")
}

func main() {
	// defer son ejecuciones que se realizan al final de la función que la rodea no import
	// en donde se coloque
	defer done()
	fmt.Println("Main: Start")
	fmt.Println("Main: End")

	fmt.Println("Using Anonymous defer function")
	defer func() {
		fmt.Println("We're anonymous, We're watching you ")
	}()
	fmt.Println("Main: Covid stars")
	fmt.Println("Main: Government said the end is nigh")
}
