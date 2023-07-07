package main

import "fmt"

func main() {
	// Array literals
	a := [3]int{1, 2, 3}
	fmt.Println(a)

	b := [...]string{"Hello", "Gophers."}
	fmt.Println(b)

	var c [2]int
	c[0] = 101
	c[1] = 1010
	fmt.Println(c[0])
	fmt.Println(c[1])

	fmt.Printf("a is type %T\n", a)
	fmt.Printf("b is type %T\n", b)
	fmt.Printf("c is type %T\n", c)
}
