package main

import (
	"fmt"
	"strconv"
)

// Function to check if a number is a palindrome
func isPalindrome(num int64) bool {
	strNum := strconv.FormatInt(num, 10)
	start, end := 0, len(strNum)-1
	for start < end {
		if strNum[start] != strNum[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// Function to find the largest palindrome from the product of two 3-digit numbers
func largestPalindrome() int64 {
	var maxPalindrome int64
	// Iterate through all pairs of 3-digit numbers
	for i := int64(999); i >= 100; i-- {
		for j := i; j >= 100; j-- {
			product := i * j
			if isPalindrome(product) && product > maxPalindrome {
				maxPalindrome = product
			}
		}
	}
	return maxPalindrome
}

func main() {
	result := largestPalindrome()
	fmt.Println("Largest palindrome made from the product of two 3-digit numbers:", result)
}
