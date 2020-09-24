package main

import "fmt"

func main() {
	fmt.Println("1. Faktor Bilangan")
	var bilangan int = 33489857205
	// fmt.Scan(&bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}

}
