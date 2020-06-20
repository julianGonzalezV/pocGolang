package main

import "fmt"

var level = "pkg"

func main() {
	fmt.Println("Main start :", level)
	if true {
		fmt.Println("Block start :", level)
		funcA()
	}

	fmt.Println("shadowed() starts", level)
	shadowed()
}
func funcA() {
	fmt.Println("funcA start :", level)
}

func shadowed() {
	fmt.Println("Main start :", level)
	// Create a shadow variable
	level := 42
	if true {
		fmt.Println("Block start :", level)
		funcA()
	}
	fmt.Println("Main end :", level)
}
