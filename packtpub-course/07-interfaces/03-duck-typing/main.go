package main

// "If it looks like a duck, swims like a duck, and quacks like a duck, then it must be a duck.

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct {
}

func main() {
	c := cat{}
	chatter(c)
}
func (c cat) Speak() string {
	return "Purr Meow"
}

/// cat matches the Speak() method of the Speaker{} interface, so a cat is a Speaker{}:
func chatter(s Speaker) {
	fmt.Println(s.Speak())
}
