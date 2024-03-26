package main

import "fmt"

type transformerFunc = func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	transformNumbers(&numbers, doubleNumber)
	fmt.Println(numbers)
	transformNumbers(&numbers, tripleNumber)
	fmt.Println(numbers)
}

func transformNumbers(numbers *[]int, transformer transformerFunc) {
	for idx, number := range *numbers {
		(*numbers)[idx] = transformer(number)
	}
}

func doubleNumber(number int) int {
	return number * 2
}

func tripleNumber(number int) int {
	return number * 3
}
