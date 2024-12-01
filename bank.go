package main

import "fmt"

func bank() {
	var accountBalance float64 = 1000
	fmt.Println("Welcom to yoyo Bank!")

	for {
		showOptions()
		choice := setUserInput("Your choice: ")

		if choice == 1 {
			fmt.Println("you current balance is: ", accountBalance)
		} else if choice == 2 {
			newDeposite := setUserInput("Deposite: ")
			if newDeposite <= 0 {
				fmt.Println("Please add deposite more then 0")
				continue
			}
			accountBalance += newDeposite
			fmt.Println("Your current balance now is: ", accountBalance)
		} else if choice == 3 {
			withdrawMoney := setUserInput("How much money you want to withdraw? ")
			if withdrawMoney > accountBalance {
				fmt.Println("you can't withdraw money more than you have")
				continue
			}
			if withdrawMoney <= 0 {
				fmt.Println("you can withdraw moeny more than 0")
				continue
			}
			accountBalance -= withdrawMoney
			fmt.Println("Your current balance now is: ", accountBalance)
		} else {
			fmt.Println("Thank you for banking with us! bye bye")
			break
		}
	}

}

func setUserInput(inputText string) float64 {
	var input float64
	fmt.Print(inputText)
	fmt.Scan(&input)
	return input
}

func showOptions() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance?")
	fmt.Println("2. Deposite money?")
	fmt.Println("3. Withdraw money?")
	fmt.Println("4. Exit")
}
