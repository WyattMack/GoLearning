package main

import "fmt"

func main() {
	xi := []int{41, 42, 43}
	fmt.Println(xi)
	fmt.Println("--------------------")

	xi = append(xi, 44, 101, 102, 103)
	fmt.Println(xi)
}
