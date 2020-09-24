package main

import (
	"fmt"
	"math"
	"time"
)

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
func isPrimeSqrt(value int) bool {
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
func main() {
	start := time.Now()

	// fmt.Println(IsPrime(1500450271))
	fmt.Println(isPrimeSqrt(33489857205))
	fmt.Println(isPrimeSqrt(96545899))
	fmt.Println(isPrimeSqrt(1000000007))

	total := time.Since(start)
	fmt.Printf("Time taken for tests: %s\n", total.String())
}
