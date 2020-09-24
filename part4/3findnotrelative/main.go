package main

import "fmt"

func notrelative(arr1 []int, arr2 []int) {
	var j int = 0
	for i := 0; i < len(arr1); i++ {
		for j < len(arr2) && arr2[j] < arr1[i] {
			// fmt.Println(arr2[j])
			j++
		}

		if arr2[j] != arr1[i] {
			fmt.Println(arr1[i])
		} else {
			j++
		}
	}
}
func main() {
	notrelative([]int{3, 6, 10, 12, 15}, []int{1, 3, 5, 10, 16})
}
