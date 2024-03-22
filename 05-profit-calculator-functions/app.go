package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Revenue: ")
	fmt.Scan(&revenue)
	outputText("Expenses: ")
	fmt.Scan(&expenses)
	outputText("Tax Rate: ")
	fmt.Scan(&taxRate)

	var profit = calculateProfit(revenue, expenses)
	// var tax = profit * (taxRate / 100)
	// var profitAfterTax = profit - tax
	var profitAfterTax = profit * (1 - (taxRate / 100))

	fmt.Println()
	fmt.Println("Profit:", profit)
	fmt.Println("Profit after tax:", profitAfterTax)
}

func outputText(text string) {
	fmt.Println(text)
}

func calculateProfit(revenue float64, expenses float64) (profit float64) {
	profit = revenue - expenses
	return profit
	// return
}
