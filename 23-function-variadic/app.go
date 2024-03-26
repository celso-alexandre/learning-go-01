package main

// Variadic function = a function that can accept a variable number of arguments.
func main() {
	println(sum(1, 2))          // 3
	println(sum(1, 2, 3, 4))    // 10
	println(sum(1, 2, 3, 4, 5)) // 15
	println(sum())              // 0
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
