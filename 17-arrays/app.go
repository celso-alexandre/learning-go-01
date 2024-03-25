package main

import "fmt"

func main() {
	prices := [4]float64{1, 2.99, 3, 4.99}
	fmt.Println(prices)

	// Slices - It references the same array in memory, but as a subset of it.
	// featuredPrices := prices[:3] // from index 0 to index < 3 (0 to 2)
	featuredPrices := prices[1:3] // from index 1 to index < 3 (1 to 2)
	fmt.Println(featuredPrices)   // [2.99 3]

	featuredPrices[0] = 999.99
	fmt.Println(prices) // [1 999.99 3 4.99]
}
