package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("x is %v\n", x)
	fmt.Printf("y is %v\n", y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are both less than four")
	} else if x > 6 && y > 6 {
		fmt.Println("x and y are both greater than six")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is four, five, or six")
	} else if y != 5 {
		fmt.Println("y is not five")
	} else {
		fmt.Println("None of the conditions were met.")
	}
}
