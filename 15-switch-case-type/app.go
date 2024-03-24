package main

import "fmt"

// func printSomething(value any) {
func printSomething(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Printf("%v is an int\n", value)
	case string:
		fmt.Printf("%v is a string\n", value)
	case float64:
		fmt.Printf("%v is a float64\n", value)
		// default:
		// 	fmt.Printf("%v is another type\n", value)
	}
}

func main() {
	printSomething("Hello")
	printSomething(1)
	printSomething(1.1)
	printSomething([]int{1, 2, 3})
}
