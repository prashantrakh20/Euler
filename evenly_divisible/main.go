package main

import "fmt"

func main() {
	// Start with 20 and keep checking
	num := 20

	// Keep increasing num until it's divisible by all numbers from 1 to 20
	for {
		// Assume the number is divisible by all numbers from 1 to 20
		isDivisible := true

		// Check divisibility for all numbers from 1 to 20
		for i := 1; i <= 20; i++ {
			if num%i != 0 {
				// If it's not divisible by i, set isDivisible to false and break
				isDivisible = false
				break
			}
		}

		// If num is divisible by all numbers, print and exit
		if isDivisible {
			fmt.Println("The smallest multiple divisible by all numbers from 1 to 20 is:", num)
			break
		}

		// If not divisible by all, increment num and try again
		num++
	}
}
