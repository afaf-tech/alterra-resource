package main

import (
	"fmt"
	"strings"
)

func try(A string, B string) string {
	var result string
	if len(A) > len(B) {
		if strings.Contains(A, B) {
			result = B
		}
	} else {
		if strings.Contains(B, A) {
			result = A
		}
	}

	return result
}
func main() {
	B := "AKASHI"
	A := "AKA"

	fmt.Println(try(A, B))

	D := "KANGAROO"
	C := "KANG"
	fmt.Println(try(C, D))

}
