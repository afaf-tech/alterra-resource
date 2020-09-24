package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("4. Exponentiation")
	var bilangan float64
	var pangkat float64
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&bilangan)
	// fmt.Println("")
	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&pangkat)

	var result = math.Pow(bilangan, pangkat)
	fmt.Println("hasil = ", result)

}
