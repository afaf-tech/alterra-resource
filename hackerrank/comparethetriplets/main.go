package main

import (
	"fmt"
)

// Complete the compareTriplets function below.
// func compareTriplets(a []int32, b []int32) []int32 {

// }

func main() {
	var a [3]int
	var b [3]int
	fmt.Scan(&a[0], &a[1], &a[2])
	fmt.Scan(&b[0], &b[1], &b[2])
	var ar, br int
	for i := 0; i < 3; i++ {
		//compare here
		if a[i] > b[i] {
			ar++
		} else if b[i] > a[i] {
			br++
		}
	}
	fmt.Println(ar, br)
}
