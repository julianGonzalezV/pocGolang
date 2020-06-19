package main

import (
	"fmt"
	"time"
)

// cuando se tenga variables relacionadas a un tema :)
var (
	valor1 string    = "valor 1"
	valor2 time.Time = time.Now()
)

func main() {
	fmt.Println(valor1, valor2)
}
