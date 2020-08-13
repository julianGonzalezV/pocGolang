package main

import "fmt"

/// Sales tax calculator
/*
1- Create a calculator that calculates the sales tax for a single item.
2- The calculator must take the items cost and its sales tax rate.
3- Sum the sales tax and print the total amount of sales tax required
  for the following items:*/

func main() {

	item1 := Item{itemName: "Cake", cost: 0.99, salesTaxR: 0.075}
	item2 := Item{itemName: "Milk", cost: 2.75, salesTaxR: 0.015}
	item3 := Item{itemName: "Butter", cost: 0.87, salesTaxR: 0.02}
	result := item1.salexTaxCalculator() + item2.salexTaxCalculator() + item3.salexTaxCalculator()
	fmt.Print("Sales Tax Total: ", result)
}

type Item struct {
	itemName  string
	cost      float32
	salesTaxR float32
}

/// al pasarle item quiere decir que va a ser una propiedad del struct Item(line 21)
func (item Item) salexTaxCalculator() float32 {
	return item.cost * item.salesTaxR
}
