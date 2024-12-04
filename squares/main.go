package main

import "fmt"

func main() {
    count := 100

    var sum_numbers, sum_squares int64

    // Loop to calculate the sum of the numbers and the sum of the squares
    for i := 1; i <= count; i++ {
        sum_squares += int64(i * i)  // sum of squares of i
        sum_numbers += int64(i)      // sum of numbers
    }

    // Calculate the square of the sum of numbers
    square_of_sum := sum_numbers * sum_numbers

    // Calculate the difference between the square of the sum and the sum of the squares
    sum_difference := square_of_sum - sum_squares

    // Print the result
    fmt.Printf("Difference is %d\n", sum_difference)
}
