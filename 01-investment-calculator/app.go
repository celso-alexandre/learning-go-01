package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1_000
	var expectedReturnRateYearly float64 = 12
	var expectedReturnRateMonthly float64 = expectedReturnRateYearly / 12
	var expectedReturnRateDecimal float64 = 1 + (expectedReturnRateMonthly / 100)
	var months float64 = 12

	// var finalAmount = investmentAmount
	// for i := 0; i < months; i++ {
	// 	finalAmount = finalAmount * (1 + expectedReturnRateDecimal)
	// }
	var finalAmount = investmentAmount * math.Pow(expectedReturnRateDecimal, months)
	fmt.Println("Final amount after", months, "months is", finalAmount)
}
