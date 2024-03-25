package main

import "fmt"

type floatMap = map[string]float64

func main() {
	// make pre-allocates the memory for a given array length. It is a optimization technique.
	names := make([]string, 2)
	// names = append(names, "celso")
	names[0] = "celso"
	// names = append(names, "alexandre")
	names[1] = "alexandre"
	fmt.Println(names)

	// for range names {
	for _, name := range names {
		fmt.Println("Hello " + name)
	}

	// courseRatings := map[string]float64{}
	courseRatings := make(floatMap, 2)
	courseRatings["go"] = 5
	courseRatings["react"] = 4.5
	fmt.Println(courseRatings)

	for course, rating := range courseRatings {
		fmt.Println("Course: " + course + ". Stars:" + fmt.Sprintf("%.2f", rating))
	}
}
