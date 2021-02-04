package main

import "fmt"

/// por idiomatic trate de usar siempre the 'er' suffix en interfaces
type Speaker interface {
	Speak() string
}
type cat struct {
	name string
	age  int
}

/// note que la implementación es IMPLICITA no se coloca en la linea 8 ...cat implements Speaker{..}
func (c cat) Speak() string {
	return "Purr Meow"
}

/// interesante: Ahpra cat implemente Stringer interface
/// este es como tener el famoso ToString en java
/// For formatting when printing values
func (c cat) String() string {
	return fmt.Sprintf("String impl %v (%v years old)", c.name, c.age)
}

/// Los types que implementan pueden tener funciones adicionales, se requiere que si deba
/// implementar tooodos los de la interface, en este caso Speaker tiene solo Speak() function
func (c cat) Greeting() {
	fmt.Println("Meow,Meow!!!!mmmeeeeoooowwww")
}

func main() {
	c := cat{name: "Michin", age: 9}
	fmt.Println(c.Speak())
	c.Greeting()
	fmt.Println(c)
}
