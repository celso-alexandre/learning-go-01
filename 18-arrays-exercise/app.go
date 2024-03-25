package main

import "fmt"

type Product struct {
	title string
	id    int
	price float64
}

func main() {
	hobbies := [3]string{"programming", "cooking", "gaming"}
	fmt.Println("Output (print) that array in the command line.", hobbies)
	fmt.Println("The first element (standalone)", hobbies[0])
	fmt.Println("The second and third element combined as a new list", hobbies[1:]) // cooking, gaming
	firstAndSecondHobbies := hobbies[:2]
	fmt.Println("Create a slice based on the first element that contains the first and second elements.", firstAndSecondHobbies)
	firstAndSecondHobbies2 := hobbies[0:2]
	fmt.Println("Create a slice based on the first element that contains the first and second elements.", firstAndSecondHobbies2) // programming, cooking
	firstAndSecondHobbies = hobbies[1:]
	fmt.Println("Re-slice the slice from (3) and change it to contain the second and last element of the original array.", firstAndSecondHobbies) // cooking, gaming

	courseGoals := []string{"Learn Go", "Learn Rust"}
	fmt.Println("Create a 'dynamic array' that contains your course goals (at least 2 goals)", courseGoals)
	courseGoals[1] = "Learn Python"
	courseGoals = append(courseGoals, "Learn Elixir")
	fmt.Println("Set the second goal to a different one AND then add a third goal to that existing dynamic array", courseGoals)

	products := []Product{
		{"Product 1", 11, 10.0},
		{"Product 2", 12, 20.0},
	}
	fmt.Println("Create a 'Product' struct with title, id, price and create a dynamic list of products (at least 2 products)", products)
	products = append(products, Product{"Product 3", 13, 30.0})
	fmt.Println("Then add a third product to the existing list of products.", products)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
