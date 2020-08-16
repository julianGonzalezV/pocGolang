package main

import (
	"fmt"
	"strings"
)

// Escrituras de este estilo se conoce como  blank identifier, para ingnorar valores
// en una asignación, pero NO SE recomienda ignorar errores por obvais razones
//_, err := file.Read(bytes)

func greeting() (name string, age int) {
	name = "John"
	age = 21
	return name, age
}

// naked return
func greetingNaked() (name string, age int) {
	name = "Wonder Woman  Naked"
	age = 2100
	return
}

func csvHdrCol(header []string) map[int]string {
	csvHeadersToColumnIndex := make(map[int]string)
	for i, v := range header {
		//fmt.Println("v ", v)
		v = strings.TrimSpace(v) // removes paces at the end of beggining
		//fmt.Println("v.TrimSpace =>", v)
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	return csvHeadersToColumnIndex
}

func main() {
	n1, age1 := greeting()
	fmt.Println("No Naked :) => ", n1, age1)
	n1, age1 = greetingNaked()
	fmt.Println("Naked Return => ", n1, age1)

	fmt.Println("Return Values")
	hdr := []string{"empid", "employee", "address", "hours worked", "hourly rate", "manager"}
	result := csvHdrCol(hdr)
	fmt.Println("Result:")
	fmt.Println(result)
	fmt.Println()
	hdr2 := []string{"employee", "empid", "hours worked", "address", "manager", "hourly rate"}
	result2 := csvHdrCol(hdr2)
	fmt.Println("Result2:")
	fmt.Println(result2)
	fmt.Println()
}
