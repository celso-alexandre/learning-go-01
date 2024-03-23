package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := readValueFromStdIo("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := readValueFromStdIo("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := readValueFromStdIo("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	profit, profitAfterTax, ratio := calculateProfit(revenue, expenses, taxRate)

	fmt.Println()
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Profit after tax: %.2f\n", profitAfterTax)
	fmt.Printf("Profit after tax ratio: %.2f\n", ratio)

	var reportFile = fmt.Sprintf("revenue,expenses,taxRate,profit,profitAfterTax,profitAfterTaxRatio\n%.2f,%.2f,%.2f,%.2f,%.2f,%.2f\n", revenue, expenses, taxRate, profit, profitAfterTax, ratio)
	os.WriteFile("report.csv", []byte(reportFile), 0644)
}

func readValueFromStdIo(label string) (float64, error) {
	var value float64
	fmt.Print(label)
	fmt.Scan(&value)
	if value < 0 {
		return value, errors.New("value cannot be negative")
	}

	return value, nil
}

func calculateProfit(revenue float64, expenses float64, taxRate float64) (profit float64, profitAfterTax float64, ratio float64) {
	profit = revenue - expenses
	profitAfterTax = profit * (1 - (taxRate / 100))
	ratio = profit / profitAfterTax
	return profit, profitAfterTax, ratio
}
