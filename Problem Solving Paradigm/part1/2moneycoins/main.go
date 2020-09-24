package main

import "fmt"

func coins(money int) []int {
	pecahan := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	temp := []int{}
	for _, pecah := range pecahan {
		for money > pecah {
			money -= pecah
			temp = append(temp, pecah)
		}
	}

	return temp
}

func main() {
	fmt.Println(coins(543))
}
