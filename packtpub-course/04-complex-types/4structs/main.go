package main

import "fmt"

type id string

func getIDs() (id, id, id) {
	var id1 id
	var id2 id = "1234-5678"
	var id3 id
	id3 = "1234-5678"
	//id4 := "otroo" // id 4 se puede usar pero OJO que no es de tipo 'id' sino que es un string.
	//por eso no se puede usar en al return
	//var id5 id = id4, tampoco funciona debe ser al contrario--ver id3
	return id1, id2, id3
}

type user struct {
	name    string
	age     int
	balance float64
	member  bool
}

func getUsers() []user {
	u1 := user{
		name:    "Tracy",
		age:     51,
		balance: 98.43,
		member:  true,
	}
	u2 := user{
		age:  19,
		name: "Nick",
	}
	u3 := user{
		"Bob",
		25,
		0,
		false,
	}
	var u4 user
	u4.name = "Sue"
	u4.age = 31
	u4.member = true
	u4.balance = 17.09
	return []user{u1, u2, u3, u4}
}

func main() {
	fmt.Println("Custom types ..asimple one")
	v1, v2, v3 := getIDs()
	fmt.Println(v1, v2, v3)

	fmt.Println(":::::::Struct types:::::::::")
	// Tipo que permite especificar sus propiedades y tipos de dato de cada uno
	// No tiene alguna forma de herencia, los designers hasta la fecha de publicado en libro
	// consideran que herancia ha causado muchos problemas
	//

	fmt.Println(getUsers())

}
