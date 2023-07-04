package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// var i int
	// for i = 0; i < 100; i++ {
	// 	fmt.Println(i)
	// }

	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		y := rand.Intn(10)
		switch {
		case x < 4 && y < 4:
			fmt.Printf("x and y are both less than four. %v, %v\n", x, y)
		case x > 6 && y > 6:
			fmt.Printf("x and y are both greater than six. %v, %v\n", x, y)
		case x >= 4 && x <= 6:
			fmt.Printf("x is between four and six (inclusively). %v\n", x)
		case y != 5:
			fmt.Printf("y is not equal to five. %v\n", y)
		default:
			fmt.Printf("None of these conditions were met. %v, %v\n", x, y)
		}
	}
}
