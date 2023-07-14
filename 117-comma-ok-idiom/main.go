package main

import "fmt"

func main() {
	n := make(map[string]int)
	n["Todd"] = 50
	n["Henry"] = 24
	n["Jen"] = 32

	for k, v := range n {
		fmt.Println(k, v)
	}
	for _, v := range n {
		fmt.Println(v)
	}
	for k := range n {
		fmt.Println(k)
	}

	v, ok := n["George"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("No value found")
	}
}
