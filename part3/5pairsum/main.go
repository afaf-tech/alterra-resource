package main

import (
	"fmt"
)

func PairSum(arr []int, target int) (int, int) {

	if len(arr) <= 1 {
		return 0, 0
	}
	storage := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		if val, ok := storage[arr[i+1]]; ok {
			return val, i + 1
		} else {
			storage[target-arr[i+1]] = i + 1
		}
	}
	return 0, 0
}

func main() {
	a, b := PairSum([]int{1, 2, 3, 4, 6}, 6)
	fmt.Println(a, b)
}
