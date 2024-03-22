package main

import "fmt"

func main() {
	revenue := readValueFromStdIo("Revenue: ")
	expenses := readValueFromStdIo("Expenses: ")
	taxRate := readValueFromStdIo("Tax Rate: ")

	profit, profitAfterTax, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Println()
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit after tax: %.2f\n", profitAfterTax)
	fmt.Printf("Profit after tax ratio: %.2f\n", ratio)
}

func readValueFromStdIo(label string) float64 {
	var value float64
	fmt.Print(label)
	fmt.Scan(&value)
	return value
}

func calculateProfit(revenue float64, expenses float64, taxRate float64) (profit float64, profitAfterTax float64, ratio float64) {
	profit = revenue - expenses
	profitAfterTax = profit * (1 - (taxRate / 100))
	ratio = profit / profitAfterTax
	return profit, profitAfterTax, ratio
	// return
}
