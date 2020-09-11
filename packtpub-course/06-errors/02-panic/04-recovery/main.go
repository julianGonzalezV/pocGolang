package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func payDay(hoursWorked, hourlyRate int) int {
	report := func() {
		// Recovery se usa para revisar el return value de la funcion actual, por eso debe ser defer (L25 	defer report())
		// ya que se ejecutaría al final
		if r := recover(); r != nil { // Si r no es nil entonces un panic ha ocurrido
			if r == ErrHourlyRate {
				fmt.Printf("hourly rate::::: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours worked:::: %d\nerr: %v\n\n", hoursWorked, r)
			}
		}
	}
	defer report()
	fmt.Printf("Pay was calculated based on:\nhours worked: %d\nhourly Rate: %d\n", hoursWorked, hourlyRate)

	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		// passing an error to the panic
		panic(ErrHoursWorked)
	}
	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := hoursWorked * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate
}

func main() {
	pay := payDay(100, 25)
	fmt.Println(pay)
	pay = payDay(100, 200)
	fmt.Println(pay)
	pay = payDay(60, 25)
	fmt.Println(pay)
}
