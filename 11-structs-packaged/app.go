package main

import (
	"fmt"

	"github.com/celso-alexandre/learning-go-01/11-structs-packaged/user"
)

func main() {
	firstName := getUserData("Enter your first name: ")
	lastName := getUserData("Enter your last name: ")
	birthDate := getUserData("Enter your birth date (MM/DD/YYYY): ")

	appUser, err := user.NewUser(
		firstName,
		lastName,
		birthDate,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.PrintUserData()
	appUser.ClearUserName()
	appUser.PrintUserData()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var input string
	fmt.Scanln(&input)
	return input
}
