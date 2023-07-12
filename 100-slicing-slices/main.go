package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 41, 42, 43, 44, 45, 101, 102, 103}
	fmt.Printf("xi - %#v\n", xi)
	fmt.Println("------------------")

	// Range is [inclusive:exclusive]
	fmt.Printf("xi - %#v\n", xi[0:6])
	fmt.Println("------------------")

	// Range is [Inclusive:]
	fmt.Printf("xi - %#v\n", xi[6:])
	fmt.Println("------------------")

	// Range is [:Exclusive]
	fmt.Printf("xi - %#v\n", xi[:5])
	fmt.Println("------------------")

	//Or you can get it all with [:], similar to the first statement
	fmt.Printf("xi - %#v\n", xi[:])
	fmt.Println("------------------")
}
