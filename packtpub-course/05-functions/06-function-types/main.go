package main

import "fmt"

type calc func(int, int) string

func calculator(f func(int, int) int, i, j int) int {
	return f(i, j)
}
func add(i, j int) int {
	return i + j
}
func subtract(i, j int) int {
	return i - j
}

func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func managerSalary(baseSalary, bonus int) int {
	return baseSalary + bonus
}
func developerSalary(hourlyRate, hoursWorked int) int {
	return hourlyRate * hoursWorked
}

func main() {
	fmt.Println("Function Type 1::::::::")
	/// Note como add es una funcion que cumple con la firma de type calc (línea 5)
	operationValue := calculator(add, 5, 6)
	fmt.Println("operationValue", operationValue)

	operationValue2 := calculator(subtract, 5, 6)
	fmt.Println("operationValue2", operationValue2)

	fmt.Println("Calculate Salary::::::::")
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)
	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)

}
