package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Printf("x - %v\n", x[:5])
	fmt.Printf("x - %v\n", x[5:])
	fmt.Printf("x - %v\n", x[2:7])
	fmt.Printf("x - %v\n", x[1:6])
}
