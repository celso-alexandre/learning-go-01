package main

import "fmt"

func main() {
	websites := map[string]string{
		"aws":    "https://aws.com",
		"google": "https://google.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["aws"])
	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)
	delete(websites, "linkedin")
	fmt.Println(websites)
}
