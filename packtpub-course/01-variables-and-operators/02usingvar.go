package main

import (
	"fmt"
)

var global1 string = "Global "

// Patient xx
type Patient struct {
	age            int
	peanultAllergy bool
	name           string
	familyName     string
}

func defPrint(p Patient) {
	fmt.Print(p)
}

func main() {
	var local1 = "Local => "
	fmt.Println(global1, local1) // Add \n
	defPrint(Patient{age: 33, peanultAllergy: false, name: "juli", familyName: "fam"})
}
