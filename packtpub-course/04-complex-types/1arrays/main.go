package main

import "fmt"

/// Use Array for managing sorted collections of data
func arraysKeys() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 0} // note que el 9 le da el size máximo la array , le estas diciendo que en la pos 9 el valor será 1
	fmt.Println("arr2 => ", arr2)
	arr3 := [10]int{1, 9: 10, 4: 5}
	return arr1 == arr2, arr1 == arr3, arr3
}

func message() string {
	arr := [...]string{
		"ready",
		"Get",
		"Go",
		"to",
	}
	// arr[4] --out of the bound
	return fmt.Sprintln(arr[1], arr[0], arr[3], arr[2])
}

// hasta acá solo sé que me tocó colocar el tamaño pero esperemos más adelante a ver
// porque en muchos casos no conocemos éste dato
func writing() [4]string {
	arr := [4]string{"a"}
	// arr[4] --out of the bound
	return arr
}

func Loop() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n", i, arr[i])
	}
	return m
}

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
	}
	return arr
}

func main() {
	var arr [10]int
	/// %#v\n to get extrainfo eg the data type
	fmt.Printf("Array initialization => %#v\n", arr) /// [0 0 0 0 0 0 0 0 0 0]

	fmt.Println("Comparing Array:")
	var arr1 [5]int
	arr2 := [5]int{3} // quiere decir que la primera posición va a tener el valor de 3
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 0}
	arr5 := [5]int{0, 0, 0, 0, 9}
	//arr6 := [9]int{0, 0, 0, 0, 0} ///arr1 == arr6 does not compile due te both are differents cause their size
	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	fmt.Println(arr1 == arr2, arr1 == arr3, arr1 == arr4, arr1 == arr5)

	fmt.Println("Arrays key:")
	comp1, comp2, arr := arraysKeys()
	fmt.Println("[10]int == [...]{9:0}       :", comp1)
	fmt.Println("[10]int == [10]int{1, 9: 10, 4: 5}}:", comp2)
	fmt.Println("arr               :", arr)

	fmt.Println("///////Reading from an Array")
	fmt.Print(message())

	fmt.Println("///////Writing to an Array")
	fmt.Print(writing())

	fmt.Println("///////Loop an Array")
	fmt.Print(Loop())

	fmt.Println("///////Modifying an Array")
	arrIn := [10]int{}
	fmt.Println("original", arrIn)
	fmt.Println(arrIn, fillArray(arrIn))
	fmt.Println(arrIn, opArray(fillArray(arrIn)))

}
