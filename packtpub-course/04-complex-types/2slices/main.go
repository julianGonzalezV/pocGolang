package main

import (
	"fmt"
	"os"
)

/// Slice son similares a array de heho es una capa sobre array pero con la diferencia
/// que no hay que procuparse por el size, adios funciones retornando aaarays indicadndo su tamanio eg
/// func x():[10]int    ..que pasa si necesito una función que aplique a cualeuier array? debo crear nmil funciones por cada tamaño?
/// Slice al rescate!
/// -Expandir a su necesidad
/// -Comun verlo en muchos desarrollos

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {
			longest = args[i]
		}
	}
	return longest
}

func appending(extraLocals []string) []string {
	var locales []string
	fmt.Println("appending")
	locales = append(locales, "en_US", "fr_FR")
	locales = append(locales, extraLocals...)
	fmt.Println("appending2")
	return locales
}

func fromSlice() string {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	m := fmt.Sprintln("First :", s[0], s[0:1], s[:1])
	m += fmt.Sprintln("Last  :", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	m += fmt.Sprintln("First 5 :", s[:5])
	m += fmt.Sprintln("Last 4 :", s[5:])
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}

func genSlices() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)
	return s1, s2, s3
}

func linkedSlices() (int, int, int) {
	fmt.Println("linkedSlices")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	fmt.Println(s1, s2, s3)
	s1[3] = 99
	fmt.Println(s1, s2, s3)
	return s1[3], s2[3], s3[3]
}

func noLinkedSlices() (int, int, int) {
	fmt.Println("noLinkedSlices")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	s3 := s1[:]
	fmt.Println(s1, s2, s3)
	s1 = append(s1, 6)
	fmt.Println(s1, s2, s3)
	s1[3] = 99
	fmt.Println(s1, s2, s3)
	return s1[3], s2[3], s3[3]
}

func capLinked() (int, int, int) {
	fmt.Println("capLinked")
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5

	s2 := s1
	s3 := s1[:]
	fmt.Println(s1, s2, s3)
	s1 = append(s1, 6)
	fmt.Println(s1, s2, s3)
	s1[3] = 99
	fmt.Println(s1, s2, s3)
	return s1[3], s2[3], s3[3]
}

func copyNoLink() (int, int, int) {
	fmt.Println("copyNoLink")
	s1 := []int{1, 2, 3, 4, 5}

	s2 := make([]int, len(s1))
	// OJO importante llamar la funcion copy ya qye sino, lo que estariamos haciendo es crear un
	// Slice s2 con el tamaño de s1 peero con ceros!!
	copied := copy(s2, s1)
	fmt.Println("copied: ", copied)
	s3 := s1[:]
	fmt.Println(s1, s2, s3)
	s1 = append(s1, 6)
	fmt.Println(s1, s2, s3)
	s1[3] = 99
	s2[3] = 77
	fmt.Println(s1, s2, s3)
	return s1[3], s2[3], s3[3]
}

func appendNoLink() (int, int, int) {
	fmt.Println("appendNoLink")
	s1 := []int{1, 2, 3, 4, 5}
	s2 := append([]int{}, s1...) // This method is the most commonly seen in real-world code when you need to ensure you get a copy of the value
	/* una fora más optima de evitar consumo de memoria para el emty slice en el append []int{}
	// trate de uasrlo en su vida diaria :)
		 s1 := []int{1, 2, 3, 4, 5}
	  s2 := append(s1[:0:0], s1...)
	*/
	s3 := s1[:]
	fmt.Println(s1, s2, s3)
	s1 = append(s1, 6)
	fmt.Println(s1, s2, s3)
	s1[3] = 99
	s2[3] = 77
	fmt.Println(s1, s2, s3)
	return s1[3], s2[3], s3[3]
}

func main() {
	if longest := findLongest(getPassedArgs(1)); len(longest) > 0 {
		fmt.Println("The longest word passed was:", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}

	fmt.Println("::::::::::Slice Appending::::::::::")
	locales := appending(getPassedArgs(1))
	fmt.Println("Locales to use:", locales)

	fmt.Println("::::::::::Slice from others slices::::::::::")
	fmt.Print(fromSlice())

	fmt.Println("::::::::::Using make to Control the Capacity of a Slice::::::::::")
	/// This is for  managing the hidden array: todo slice tiene su hidden array
	/// si este llega al limite entonces se crea otro más grande se aactualiza elpointer
	/// del slice hacia el nuevo array
	s1, s2, s3 := genSlices()
	fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	fmt.Printf("s2: len = %v cap = %v\n", len(s2), cap(s2))
	fmt.Printf("s3: len = %v cap = %v\n", len(s3), cap(s3))

	fmt.Println("::::::::::Controlling Internal Slice Behavior::::::::::")
	/// dado el statement del hidden array entonces no se podrán comparar 2 slices
	/// ya que es posible que sean diferente en su size del hiddden array ...solo se
	// pueden comparar contra nil, slice es un constructor especial, no es un valor ni pointer
	//:::::::copy slices
	/// Note que acá si se ve afectado el cambio, no como el ejercicio con Arrays
	/// ya que comparte ne mismo hidden array
	linkedSlices()
	// acá no pasa,osea que append es una opeación mas segura?R/ si , ya que lo que hace es crear un
	// nuevo hidden array, solo para que slice que lo alteró en su tamaño
	noLinkedSlices()

	// Acá tambien los afecta todos osea que no solo por ser append sea una operación segura
	// También depende del metodo constructor que se use, en este caso es un make que a pesar de contener
	// 5 elementos soporta(capacidad) hasta 10 items, por ende el append NO CREÓ un hidden array
	capLinked()

	// Acá S2 tiene su propio hidden array por eso no se ve afectado por la operaciones sobre s1
	// s3 quedó apuntando al old jeje de s1
	copyNoLink()

	// append([]int{}, s1...) creó otro
	appendNoLink()

}
