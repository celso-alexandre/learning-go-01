package main

import "fmt"

func main() {
	// make pre-allocates the memory for a given array length. It is a optimization technique.
	names := make([]string, 2)
	// names = append(names, "celso")
	names[0] = "celso"
	// names = append(names, "alexandre")
	names[1] = "alexandre"
	fmt.Println(names)

	// courseRatings := map[string]float64{}
	courseRatings := make(map[string]float64, 2)
	courseRatings["go"] = 5
	courseRatings["react"] = 4.5
	fmt.Println(courseRatings)
}
