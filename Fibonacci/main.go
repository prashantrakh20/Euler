
// Each new term in the Fibonacci sequence is generated by adding the previous two terms.
// By starting with 1 and 2, the first few terms will be: 1, 2, 3, 5, 8, 13, 21, ...
// By considering the terms in the Fibonacci sequence whose values do not exceed four million,
// this program calculates the sum of the even-valued terms.

package main

import "fmt"

func main() {
   a,b := 0,1
   sum := 0

   for b <= 4000000 {
     if b % 2 == 0 {
		sum += b
	  }
      a, b = b, a + b
   }
   fmt.Println("Sum of even Fibonacci numbers below 4 million:", sum)



}
