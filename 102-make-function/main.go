package main

import "fmt"

func main() {
	// first arg is defining type, second is how many elements it should be initialized with
	// and third being the capacity that should exist at maximum
	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println(len(xi))
	// cap being the capacity of the slice
	fmt.Println(cap(xi))
	fmt.Println("--------------")

	xi = append(xi, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("--------------")

	xi = append(xi, 10, 11, 12, 13, 14, 15)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
	fmt.Printf("xi - %#v\n", xi)
}
