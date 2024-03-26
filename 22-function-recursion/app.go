package main

func main() {
	println(factorial(3)) // 3 * 2 * 1 = 6
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
