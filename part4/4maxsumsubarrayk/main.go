package main

import "fmt"

func maxSumContiguous(numbers []int, size int) int {
	next := 0
	max := 0

	for _, value := range numbers[0:size] {
		max += value
	}
	next = max
	for index, value := range numbers[size:] {
		next = next - numbers[index] + value
		if next > max {
			max = next
		}
	}

	fmt.Println(next)

	return max
}
func main() {
	// find max sum with size of subarray
	// fmt.Println(maxSumContiguous([]int{2, 3, 4, 1, 5}, 2))    //7
	fmt.Println(maxSumContiguous([]int{2, 1, 5, 1, 3, 2}, 3)) //9
}
