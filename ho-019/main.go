package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(251)
	fmt.Printf("Variable x is %v\n", x)
	switch {
	case x <= 100:
		fmt.Printf("x is less than or equal to 100. It is %v\n", x)
	case x >= 101 && x < 201:
		fmt.Printf("x is between 101 and 200. It is %v\n", x)
	case x >= 201 && x < 251:
		fmt.Printf("x is between 201 and 250. It is %v\n", x)
	default:
		fmt.Println("Number is out of bounds.")
	}
	// if x <= 100 {
	// 	fmt.Printf("x is less than or equal to 100. It is %v\n", x)
	// } else if x == 101 && x < 201 {
	// 	fmt.Printf("x is between 101 and 200. It is %v\n", x)
	// } else {
	// 	fmt.Printf("x is between 201 and 250. It is %v\n", x)
	// }
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
	// fmt.Println(rand.Intn(3))
}

func init() {
	fmt.Println("This is where initalization for my program occurs. Also niladic is a function that requires \nzero arguements")
}
