package main

import "fmt"

type user struct {
	firstName string
	lastName  string
	birthDate string
}

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	printUserData(user{firstName, lastName, birthDate})
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
}
