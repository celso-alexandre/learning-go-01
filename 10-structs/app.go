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
	printUserData(appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}

func printUserData(u user) {
	fmt.Println("First name:", u.firstName)
	fmt.Println("Last name:", u.lastName)
	fmt.Println("Birth date:", u.birthDate)
	fmt.Println("Created at:", u.createdAt)
}
