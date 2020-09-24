package main

import (
	"fmt"
)

func main() {
	countries := [2]string{"Indonesia", "Singapure"}
	var even [5]int = [5]int{3, 43, 4534, 232, 232}
	// countries[0] = "Indonesia"
	// countries[1] = "Singapure"

	fas := []int{2, 1, 3, 42, 323, 32, 323, 23}
	fmt.Println(fas)
	fas = append(fas, 999)
	fmt.Println(countries)
	fmt.Println(even)
	fmt.Println(fas)
}
