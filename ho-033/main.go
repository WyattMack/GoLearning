package main

import (
	"fmt"
)

func main() {
	i := 0
	for i = 0; i < 101; i++ {
		if i%2 == 1 {
			fmt.Printf("%v is odd\n", i)
		}

	}
}
