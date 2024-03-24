package main

func main() {
	println(sum(1, 2))
	println(sum(1.5, 1.5))
	println(sum(1, 1.5))
	println(sum("Hello, ", "World!"))
}

func sum[T int | float64 | string](n1 T, n2 T) T {
	return n1 + n2
}
