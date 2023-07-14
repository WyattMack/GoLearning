package main

import "fmt"

func main() {
	a := 0
	fmt.Println(a)
	a++
	fmt.Println(a)

	fmt.Println("-------")

	m := make(map[string]int)
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])
	m["beautiful"]++
	fmt.Println(m["beautiful"])

	/*
		Actual code for this lesson skipped a lot, so this is just the demo version
		incrementing a value in a map
	*/
}
