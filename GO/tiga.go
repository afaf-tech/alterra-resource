package main

import "fmt"

func cetakTabelPerkalian(number int) {
	if !((number >= 1) && (number <= 30)) {
		fmt.Print("Invalid Number!")
	} else {
		for i := 1; i <= number; i++ {
			for j := 1; j <= number; j++ {
				fmt.Print(i*j, " ")
			}
			fmt.Print("\n")
		}
	}

}

func mains() {
	fmt.Println("Soal 3")
	// number is beetween 1 - 30
	cetakTabelPerkalian(9)
}
