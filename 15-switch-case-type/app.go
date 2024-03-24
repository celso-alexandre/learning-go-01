package main

import "fmt"

// func printSomething(value any) {
func printSomething(value interface{}) {
	intVal, ok := value.(int)
	fmt.Println(intVal, ok)

	if ok {
		fmt.Println("Value is an int", intVal)
		return
	}

	floatVal, ok := value.(float64)
	fmt.Println(floatVal, ok)

	if ok {
		fmt.Println("Value is a float", floatVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Printf("%v is an int\n", value)
	// case string:
	// 	fmt.Printf("%v is a string\n", value)
	// case float64:
	// 	fmt.Printf("%v is a float64\n", value)
	// 	// default:
	// 	// 	fmt.Printf("%v is another type\n", value)
	// }
}

func main() {
	printSomething("Hello")
	printSomething(1)
	printSomething(1.1)
	printSomething([]int{1, 2, 3})
}
