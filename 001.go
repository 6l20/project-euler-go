package main

import "fmt"

// SumOf3And5MultsGivenLimit
// Find the sum of all the multiples of 3 or 5 below given upperValue
func SumOf3And5MultsGivenLimit(upperValue int) int {
	sum := 0

	for i := 0; i < upperValue; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}

	return sum
}

func main() {
	upperValue := 10 // Should return 23
	fmt.Printf("Sum Of Multiples of 3 and 5 with upperValue %d is %d", upperValue, SumOf3And5MultsGivenLimit(upperValue))
}
