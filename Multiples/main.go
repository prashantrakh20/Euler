// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6, and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below a given limit.

package main
import "fmt"

func main (){
  sum := 0
  count := 1000
  for i := 0; i < count; i++ {
	if i%3 == 0 || i%5 == 0 {
      sum = sum + i	
	}
  }
  fmt.Println(sum)


}

