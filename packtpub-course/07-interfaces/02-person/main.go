package main

import (
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	IsMarried bool
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years old).\nMarried status: %v ", p.Name, p.Age, p.IsMarried)
}

func (p Person) Speak() string {
	return "Hi my name is: " + p.Name
}

func main() {
	p := Person{Name: "Cailyn", Age: 44, IsMarried: false}
	fmt.Println(p.Speak())
	fmt.Println(p)
}
