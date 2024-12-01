package main

import "fmt"

func profiteCalculator() {

	var adding = someCalculation(3, 4)

	revenue := setUserInput("Revenue: ")
	expenses := setUserInput("Expenses: ")
	tax_rate := setUserInput("Tax rate: ")

	fmt.Println("Your value has been stored and some bonus", adding)
	fmt.Println("Revenue", revenue)
	fmt.Println("Your Expenses", expenses)
	fmt.Println("Your tax rate", tax_rate)
}

func someCalculation(val1 int, val2 int) (value int) {
	value = val1 * val2
	return value
}
