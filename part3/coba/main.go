package main

import "fmt"

func power(numb int, pangkat int) int {
	result := 1

	for i := 0; i < pangkat; i++ {
		result = result * numb
	}

	return result
}

func rekur(num int, pangkat int) int {
	if pangkat == 0 {
		return 1
	} else if pangkat == 1 {
		return num
	} else {
		return num * rekur(num, pangkat-1)
	}
}

func simpleExp(num int, pangkat int) int {
	temp := 1
	/* Bayangkan 2^6 == 2^3 * 2^3 */
	for i := 1; i <= pangkat/2; i++ {
		temp *= num * num
	}

	if pangkat%2 != 0 {
		temp *= num
	}
	return temp
}

func main() {
	fmt.Println(power(2, 7))
	fmt.Println(rekur(2, 7))
	fmt.Println(simpleExp(2, 10))
}
