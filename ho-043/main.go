package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	/*
		Understanding that the instructor would likely want the same
		information included as the previous hands-on exercise, I've
		included the index in my solution
	*/
	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, i)
	}
	/*
		Instructor used \t to make a tab instead of a hyphen as in the
		previous example. That's the only real difference between ours.
		He then included the following to show the type of slice before
		the values.
	*/
	fmt.Printf("%T \t %#v", x, x)
}
