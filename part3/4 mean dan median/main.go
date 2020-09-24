package main

import (
	"fmt"
	"time"
)

func MeanMedian(arr []float64) (float64, float64) {
	totalmean := 0.0
	median := 0.0
	totalnumber := len(arr)

	for _, v := range arr {
		totalmean += v
	}
	var mean float64 = totalmean / float64(len(arr))

	// execute median
	if totalnumber%2 == 0 {
		median = (arr[totalnumber/2] + arr[(totalnumber/2)-1]) / 2
	} else {
		median = arr[(totalnumber-1)/2]
	}

	return mean, median
}
func main() {
	start := time.Now()

	a, b := MeanMedian([]float64{3, 5, 7, 5, 3})

	var result = fmt.Sprintf("Mean : %.1f, Median :%.1f", a, b)

	fmt.Println(result)

	total := time.Since(start)
	fmt.Printf("Time taken for tests: %s\n", total.String())

}
