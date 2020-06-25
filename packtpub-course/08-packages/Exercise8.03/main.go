package main

import "fmt"

// Executing Multiple init() Functions::::::::::::::::::::::
// super ganador par cuando desee agrupar las inicializaciones por funcionalidad
// ejemplo , conexiones de BDs, inicializacion de env variables , you choose

var name = "Gopher"

func init() {
	fmt.Println("Hello, ", name)
}
func init() {
	fmt.Println("Second")
}
func init() {
	fmt.Println("Third")
}
func main() {
	fmt.Println("Hello, main function")
}
