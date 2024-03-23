package main

import (
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

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	appUser := user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}
	// appUser.printUserData(&appUser)
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
