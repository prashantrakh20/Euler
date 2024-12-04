package main

import (
	"fmt"
	"math"
)


func primeFactors(n int) []int {
	factors := []int{}

	for n%2 == 0 {
		factors = append(factors, 2)
		n = n/2
	}

	for i := 3; i <= int(math.Sqrt(float64(n)))+1; i += 2 {
		for n%i == 0 {
			factors = append(factors, i)
			n = n/i
		}
	}

	// If n is a prime number greater than 2
	if n > 2 {
		factors = append(factors, n)
	}

	return factors
}

func largestPrimeFactor(n int) int {
	factors := primeFactors(n)
	return factors[len(factors)-1]
}

func main() {
	var number int = 6857

	factors := primeFactors(number)
    for _, v := range factors {
		fmt.Printf("%d\n", v)
	}

	largest := largestPrimeFactor(number)
	fmt.Printf("The largest prime factor of %d is: %d\n", number, largest)
	
}


