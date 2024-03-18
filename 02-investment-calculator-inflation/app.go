package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRateYearly float64 = 10
	const inflationRateMonthly float64 = (inflationRateYearly) / 12
	const inflationRateMonthlyDecimal float64 = 1 + (inflationRateMonthly / 100)

	var investmentAmount float64 = 100
	var expectedReturnRateYearly float64 = 12
	var expectedReturnRateMonthly float64 = expectedReturnRateYearly / 12
	var expectedReturnRateDecimal float64 = 1 + (expectedReturnRateMonthly / 100)
	var months float64 = 12

	var futureValue = investmentAmount * math.Pow(expectedReturnRateDecimal, months)
	var inflationInPeriod = investmentAmount * math.Pow(inflationRateMonthlyDecimal, months)
	var futureRealValue = futureValue - inflationInPeriod

	fmt.Println("Final amount after", months, "months is", futureValue)
	fmt.Println("Final real amount after", months, "months is", futureRealValue)
}
