package main

import (
	"fmt"
)

// Lo podemos tener como constantes y e valido
const (
	Sunday1  = 0
	Monday1  = 1
	Tuesday1 = 2
	//newDay = 3 esto hace que debamos modificar todo para abajo, vaya a la version iota
	Wednesday1 = 3
	Thursday1  = 4
	Friday1    = 5
	Saturday1  = 6
)

const (
	Sunday = iota
	Monday
	Tuesday
	NewDay // note como de una lo modifica
	Wednesday
	Thursday
	Friday
	Saturday
)

//Enums are a way of defining a fixed list of values that are all related

func main() {
	fmt.Println("NewDay => ", NewDay, "Wednesday => ", Wednesday)

}
