package main

import "fmt"

func main() {
	m := map[string]int{
		"Todd":  42,
		"Henry": 16,
	}

	// or

	n := make(map[string]int)
	n["Todd"] = 50
	n["Henry"] = 24

	fmt.Println("Todd's age is", m["Todd"])
	fmt.Println("Henry's age is", n["Henry"])
	fmt.Println(m)
	fmt.Println(n)
}
