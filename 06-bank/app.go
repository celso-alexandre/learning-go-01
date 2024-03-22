package main

import "fmt"

func main() {
	var balance float64
	for {
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

		switch choice {
		case 1:
			fmt.Print("Enter the amount to deposit: ")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit < 0 {
				fmt.Println("Invalid amount!")
				continue
			}
			balance += deposit
			fmt.Printf("Deposited %.2f. Current balance: %.2f\n", deposit, balance)
		case 2:
			fmt.Print("Enter the amount to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid amount!")
				continue
			}
			if withdraw > balance {
				fmt.Println("Insufficient balance!")
			} else {
				balance -= withdraw
				fmt.Printf("Withdrawn %.2f. Current balance: %.2f\n", withdraw, balance)
			}
		case 3:
			fmt.Printf("Current balance: %.2f\n", balance)
		case 4:
			fmt.Println("\nGoodbye! ( ͡~ ͜ʖ ͡°)")
			return
		}
	}
}
