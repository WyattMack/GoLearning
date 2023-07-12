package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 55, 56, 101, 102, 103, 104}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("---------------")

	xi = append(xi[:4], xi[5:]...)
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("---------------")
}
