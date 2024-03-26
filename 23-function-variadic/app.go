package main

// Variadic function = a function that can accept a variable number of arguments.
func main() {
	println(sum(1, 2))          // 3
	println(sum(1, 2, 3, 4))    // 10
	println(sum(1, 2, 3, 4, 5)) // 15
	println(sum())              // 0
}

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
