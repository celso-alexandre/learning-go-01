package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRateYearly float64 = 10
	const inflationRateMonthly float64 = (inflationRateYearly) / 12
	const inflationRateMonthlyDecimal float64 = 1 + (inflationRateMonthly / 100)

	var months float64 = 12
	var investmentAmount float64 = 100
	var expectedReturnRateYearly float64 = 12

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Expected return rate (yearly): ")
	fmt.Scan(&expectedReturnRateYearly)
	fmt.Print("Months: ")
	fmt.Scan(&months)

	var expectedReturnRateMonthly float64 = expectedReturnRateYearly / 12
	var expectedReturnRateDecimal float64 = 1 + (expectedReturnRateMonthly / 100)

	var futureValue = investmentAmount * math.Pow(expectedReturnRateDecimal, months)
	var inflationInPeriod = investmentAmount * math.Pow(inflationRateMonthlyDecimal, months)
	var futureRealValue = futureValue - inflationInPeriod

	// fmt.Println("Final amount after", months, "months is", futureValue)
	// fmt.Println("Final real amount after", months, "months is", futureRealValue)

	// tip: fmt.Sprintf does not print to console, it returns a string
	fmt.Printf("Final amount after %.0f months is %.2f\n", months, futureValue)
	// fmt.Printf(`Final amount after %.0f months is %.2f
	//`, months, futureValue)
	fmt.Printf("Final real amount after %.0f months is %.2f\n", months, futureRealValue)
}
