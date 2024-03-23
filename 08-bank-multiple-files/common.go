package main

import "fmt"

func displayAndReadChoice() uint8 {
	fmt.Println()
	fmt.Println("Welcome to the Bank!")
	fmt.Println("What do you want, huh?")
	fmt.Println("1. Deposit ( ͡° ͜ʖ ͡°)")
	fmt.Println("2. Withdraw ( ಠ ͜ʖಠ)")
	fmt.Println("3. Check Balance ( ͡~ ͜ʖ ͡°)")
	fmt.Println("4. Exit (╯ ͠° ͟ʖ ͡°)╯┻━┻")

	var choice uint8
	fmt.Print("Enter your choice: ")
	fmt.Scan(&choice)

	return choice
}
