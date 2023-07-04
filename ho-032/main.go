package main

import "fmt"

var y int

func main() {
	for {
		fmt.Printf("y is equal to %v\n", y)
		if y > 20 {
			break
		}
		y++
	}
}

// This evaluates to 'y is equal to 21' because of the formatting. The print statement runs before the
// the evaluation of the 'if' statement, and well before the break. The latent y++ allows the code to process
//  beyond 20 because once things hit 20, it has already gone beyond and looped back over.
