package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      42,
		"Moneypenny": 32,
	}

	age1 := m["James"]
	fmt.Println("James' age is ", age1)

	age2 := m["Q"]
	fmt.Println("Q's age is ", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("Q doesn't exist. Here's the zero value of an int:", v)
	}
}
