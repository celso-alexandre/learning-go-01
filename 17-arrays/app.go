package main

import "fmt"

func main() {
	prices := [4]float64{1, 2.99, 3, 4.99}
	fmt.Println(prices)

	featuredPrices := prices[1:3] // from index 1 to index < 3 (1 to 2)
	fmt.Println(featuredPrices)   // [2.99 3]
}
