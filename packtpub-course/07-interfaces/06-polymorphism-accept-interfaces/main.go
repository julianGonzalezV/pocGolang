// Tip: accepting interfaces for your APIs
// (functions, methods, and so on) and the return to be structs or concrete types

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

func main() {
	s := `{"Name":"Joe","Age":18}`
	s2 := `{"Name":"Jane","Age":21}`
	/// Enviando interfaces solo es asegurar que lo que se envían cumple con la
	// firma que la interface estblece , en este caso un arreglo de bytes
	p, err := loadPerson(strings.NewReader(s))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	p2, err := loadPerson2(s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p2)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

/// No cumple con " "accept interfaces, return structs"."
func loadPerson2(s string) (Person, error) {
	var p Person
	err := json.NewDecoder(strings.NewReader(s)).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}

/// loadPerson cumple con  "accept interfaces, return structs".
/// Note que Reader es una interface y se devuelve el Person concret
/// asi eso pordrá recibir strings.newReader, http.Request.Body, os.File
/// TIp:
/// 	Check to see whether there is an interface in the Go standard packages.
///		This will increase the number of different implementations that your
/// 	function can provide.
/// 	:o :o :o :o :o
func loadPerson(r io.Reader) (Person, error) {
	var p Person
	err := json.NewDecoder(r).Decode(&p)
	if err != nil {
		return p, err
	}
	return p, nil
}
