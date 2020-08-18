package main

/// Un empty interface no tiene métodos ..ESTO QUIERE decir que todo tipo en en Go
/// implementa una interface empty

import (
	"fmt"
)

type cat struct {
	name string
}

func main() {
	c := cat{name: "oreo"}
	i := []interface{}{42, "The book club", true, c}
	typeExample(i)

	//::::::::::::::::::::::.
	fmt.Println(":::::::::::::::::::::: Analyzing Empty interface{} Data ::::::::")
	m := make(map[string]interface{})
	a := animal{name: "oreo", category: "cat"}
	p := person{lastName: "Doe", isMarried: false, age: 19}
	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "Smith"
	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}
	for _, v := range rs {
		fmt.Println("Key: ", v.key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}
}
func typeExample(i []interface{}) {
	for _, x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is int\n", v)
		case string:
			fmt.Printf("%v is a string\n", v)
		case bool:
			fmt.Printf("a bool %v\n", v)
		case cat:
			fmt.Printf("It is a Cat %v\n", v)
		default:
			fmt.Printf("Unknown type %T\n", v)
		}
	}
}

type record struct {
	key       string
	valueType string
	data      interface{}
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}

type animal struct {
	name     string
	category string
}

func newRecord(key string, i interface{}) record {
	r := record{}
	r.key = key
	switch v := i.(type) {
	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unknown"
		r.data = v
		return r
	}
}
