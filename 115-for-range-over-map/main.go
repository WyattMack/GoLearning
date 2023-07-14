package main

import "fmt"

func main() {
	// m := map[string]int{
	// 	"Todd":  42,
	// 	"Henry": 16,
	// }

	// or

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
}
