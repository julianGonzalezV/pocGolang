package main

import (
	"fmt"
)

func main() {
	offset := 5
	fmt.Println(offset)
	offset = 10 // si usa :=  no compila porque es para creación mas no para acualización
	fmt.Println(offset)
}
