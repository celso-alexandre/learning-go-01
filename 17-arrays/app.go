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

	fmt.Println(len(featuredPrices)) // 2
	fmt.Println(cap(featuredPrices)) // 3

	// Append
	// prices = append(prices, 5.99) // NO! Append only works with slices
	pricesSlice := []float64{1, 2, 3}
	pricesSlice = append(pricesSlice, 4, 5, 6) // append creates a new slice without mutating the original one
	fmt.Println(pricesSlice)                   // [1 2 3 4 5 6]
}
