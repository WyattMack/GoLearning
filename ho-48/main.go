package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	mm := []string{"Miss", "Moneypenny", "008"}

	fmt.Println(jb)
	fmt.Println(mm)
	fmt.Println("----------------")

	//xxp being slice of slice of people
	xxp := [][]string{jb, mm}

	for i, v := range xxp {
		fmt.Println(i, v)
		for x, y := range v {
			fmt.Println(x, y)
		}
	}
}
