package main

import (
	"fmt"

	"github.com/celso-alexandre/learning-go-01/06-bank-multiple-files/fileOperations"
)

const balanceFilename = "balance"

func main() {
	balance, err := fileOperations.ReadFloatFromFile(balanceFilename)
	if err != nil {
		fileOperations.WriteFloatToFile(balanceFilename, balance)
		fmt.Println(err, ". Balance file created")
		// panic("Error reading balance file.")
	}

	fmt.Printf("Initial balance: %.2f\n", balance)

	for {
		choice := displayAndReadChoice()

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

		fileOperations.WriteFloatToFile(balanceFilename, balance)
	}
}
