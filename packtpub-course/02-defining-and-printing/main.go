package main

import (
	"fmt"
)

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
	defPrint(Patient{age: 33, peanultAllergy: false, name: "juli", familyName: "fam"})
}
