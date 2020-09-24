package main

import (
	"fmt"
)

func duplicationRemoval(nums []int) int {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	fmt.Println(duplicationRemoval([]int{2, 3, 3, 3, 6, 9, 9}))
}
