package main

import (
	"fmt"

	"seanThakur.com/app/user"
)

func newStruck() {
	firstName := setNewUserInput("Please enter your firstname: ")
	lastName := setNewUserInput("Please enter your lastname: ")
	dob := setNewUserInput("Your DOB: ")
	// user var declare with pointer
	userPointer, err := user.New(firstName, lastName, dob)
	if err != nil {
		fmt.Println(err)
		return
	}
	userPointer.GetDetails()
	userPointer.ClearUserName()
	userPointer.GetDetails()
}

func setNewUserInput(inputText string) string {
	fmt.Print(inputText)
	var userInput string
	fmt.Scanln(&userInput)
	return userInput
}
