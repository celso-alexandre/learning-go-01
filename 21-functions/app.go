package main

import "fmt"

type transformFn = func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// transformNumbers(&numbers, doubleNumber)
	transformNumbers(&numbers, func(number int) int { return number * 2 })
	fmt.Println(numbers)
	// transformNumbers(&numbers, tripleNumber)
	transformNumbers(&numbers, func(number int) int { return number * 3 })
	fmt.Println(numbers)
}

func transformNumbers(numbers *[]int, transformer transformFn) {
	for idx, number := range *numbers {
		(*numbers)[idx] = transformer(number)
	}
}

// func doubleNumber(number int) int {
// 	return number * 2
// }

// func tripleNumber(number int) int {
// 	return number * 3
// }
