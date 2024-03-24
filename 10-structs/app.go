package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) printUserData() {
	fmt.Println()
	fmt.Println("First name:", u.firstName)
	fmt.Println("Last name:", u.lastName)
	fmt.Println("Birth date:", u.birthDate)
	fmt.Println("Created at:", u.createdAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birth date are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	// appUser := user{
	// 	firstName: firstName,
	// 	lastName:  lastName,
	// 	birthDate: birthDate,
	// 	createdAt: time.Now(),
	// }
	// appUser.printUserData(&appUser)
	appUser, err := newUser(
		firstName,
		lastName,
		birthDate,
	)
	if err != nil {
		// panic(err)
		fmt.Println(err)
		return
	}
	appUser.printUserData()
	appUser.clearUserName()
	appUser.printUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
