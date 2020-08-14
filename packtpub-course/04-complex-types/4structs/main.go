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
type client struct {
	name string
	age  int
}

type contact struct {
	address string
	country string
}

type clientContact struct {
	client
	contact
}
type point struct {
	x int
	y int
}

func compare() (bool, bool) {
	point1 := struct {
		x int
		y int
	}{
		10,
		10,
	}

	point2 := struct {
		x int
		y int
	}{}
	point2.x = 10
	point2.y = 5
	point3 := point{10, 10}

	return point1 == point2, point1 == point3
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

func getClientContact() []clientContact {
	var client1 clientContact

	client2 := clientContact{}
	client2.name = "Juli"
	client2.age = 33
	client2.address = "cll 20"
	client2.country = "COLO"

	client3 := clientContact{
		client: client{
			name: "Mati",
			age:  1,
		},
		contact: contact{
			address: "cll 83",
			country: "COLO",
		},
	}

	client4 := clientContact{}
	client4.client.name = "Paola"
	client4.client.age = 101
	client4.contact.address = "cll 31"
	client4.contact.country = "COL"

	return []clientContact{client1, client2, client3, client4}

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

	fmt.Println("::::::Comparing Structs:::::::::")
	/// OJO: si contiene un slice en el struct no le va a funcionar, hidden array(ver folder 2slices)
	a, b := compare()
	fmt.Println("point1 == point2:", a)
	fmt.Println("point1 == point3:", b)

	fmt.Println("::::::Struct Embedding and Initialization:::::::::")
	clients := getClientContact()
	for i := 0; i < len(clients); i++ {
		fmt.Println(clients[i])
	}

}
