package main

import (
	"fmt"
	"math"
)

func main() {

	var size int
	fmt.Scan(&size)

	a := make([][]int, size)

	leftSum, reightSum := 0, 0

	for i := 0; i < size; i++ {
		a[i] = make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Scan(&a[i][j])
			if i == j {
				leftSum += a[i][j]
			}
			if i+j == size-1 {
				reightSum += a[i][j]
			}
		}
	}

	diff := math.Abs(float64(leftSum - reightSum))
	fmt.Println(diff)

}
