package main

import (
	"fmt"
)

type Speaker interface {
	Speak() string
}
type cat struct {
}
type dog struct {
}
type person struct {
	name string
}

func (c cat) Speak() string {
	return "Purr Meow"
}
func (d dog) Speak() string {
	return "Woof Woof"
}
func (p person) Speak() string {
	return "Hi my name is " + p.name + "."
}

/*
func catSpeak(c cat) {
	fmt.Println(c.Speak())
}
func dogSpeak(d dog) {
	fmt.Println(d.Speak())
}
func personSpeak(p person) {
	fmt.Println(p.Speak())
}*/

func saySomething(say ...Speaker) {
	for _, s := range say {
		fmt.Println(s.Speak())
	}
}

func main() {
	c := cat{}
	d := dog{}
	p := person{name: "Heather"}
	/*catSpeak(c)
	  dogSpeak(d)
	  personSpeak(p) */
	saySomething(c, d, p)
}
