package user

import (
	"errors"
	"fmt"
	"time"
)

type UserType struct {
	firstName string
	lastName  string
	dob       string
	createdAt time.Time
}

// usertype functions
func (user *UserType) GetDetails() {
	fmt.Println("Your CREDs: ", user.firstName, user.lastName, user.dob)
}

func (user *UserType) ClearUserName() {
	user.firstName = ""
	user.lastName = ""
}

// user constructor with pointer for memeory efficiency
func New(firstName, lastName, dob string) (*UserType, error) {
	if firstName == "" || lastName == "" || dob == "" {
		return nil, errors.New("all the fields should be field")
	}
	return &UserType{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
		createdAt: time.Now(),
	}, nil
}
