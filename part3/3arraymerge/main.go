package main

import "fmt"

func ArrayMerge(arrayA, arrayB []string) []string {

	isSame := make(map[string]bool)
	result := []string{}

	for _, isi := range arrayB {
		arrayA = append(arrayA, isi)
	}

	// append to map will replace the same key
	for _, isi := range arrayA {
		if _, val := isSame[isi]; !val {
			result = append(result, isi)
		}
	}

	return result
}
func main() {
	var array1 = []string{"kazuya", "jin", "lee"}
	var array2 = []string{"kazuya", "feng"}
	fmt.Println(ArrayMerge(array1, array2))
}
