package main

import (
	"fmt"
	forma "pocGolang/interface-example/interface" // esta forma es cuando el folder "interface" no se llama igual al package "interfaces"
	"pocGolang/interface-example/interfaceimp"    //note que acá n ofué necesario porque el folder y el pck se llaman igual
)

func measure(g forma.Geometry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}

func main() {
	r := interfaceimp.Rectangle{Width: 3, Height: 4}
	c := interfaceimp.Circle{Radius: 5}
	measure(r)
	measure(c)
}
