package main

import (
	"errors"
	"fmt"
	"os"
)

const balanceFile = "balance"

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile, []byte(balanceText), 0644)
}

func readBalanceFromFile() (float64, error) {
	balanceText, err := os.ReadFile(balanceFile)
	if err != nil {
		return 0, errors.New("balance file not found")
	}
	var balance float64
	_, parsingError := fmt.Sscan(string(balanceText), &balance)
	if parsingError != nil {
		return 0, errors.New("balance file is invalid")
	}
	return balance, nil
}
