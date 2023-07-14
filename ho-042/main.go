package main

import "fmt"

func main() {
	x := [5]int{1, 3, 5, 7, 9}

	for _, v := range x {
		fmt.Printf("%v - %T\n", v, v)
	}

	fmt.Println("---------------")

	//  Instructors solution, updated to y instead of x:
	y := [5]int{}

	// assign values to each index position
	for i := 0; i < 5; i++ {
		x[i] = i
	}

	// print out
	for i, v := range y {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
}
