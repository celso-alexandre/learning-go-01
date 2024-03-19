package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	var profit = revenue - expenses
	// var tax = profit * (taxRate / 100)
	// var profitAfterTax = profit - tax
	var profitAfterTax = profit * (1 - (taxRate / 100))

	fmt.Println()
	fmt.Println("Profit:", profit)
	fmt.Println("Profit after tax:", profitAfterTax)
}
