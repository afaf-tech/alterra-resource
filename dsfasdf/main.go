package main

import "fmt"

func main() {
	str := "KOTAK"
	str1 := ""
	for i := len(str) - 1; i >= 0; i-- {
		str1 += string(str[i])

	}
	if str == str1 {
		fmt.Println("Palindrome")
	} else {
		fmt.Println("Not Palindrome")
	}
	// fmt.Println(str1)
}
