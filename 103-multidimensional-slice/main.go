package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Martini", "Chocolate"}
	jm := []string{"Jenny", "Moneypenny", "Beer", "Rocky Road"}
	q := []string{"Q", "", "Capri Sun", "It's a mystery"}

	fmt.Println(jb)
	fmt.Println(jm)
	fmt.Println(q)

	// xxs being slice of slice of string.
	xxs := [][]string{jb, jm, q}
	fmt.Println(xxs)
}
