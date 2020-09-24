package main

import "fmt"

func drawSymbolicColsBox(high int) string {
	var temp string = ""
	var counter int = 0

	for i := 1; i <= high; i++ {
		for j := 0; j < high; j++ {
			counter += 1
			if counter%3 == 0 {
				temp += "* "
			} else {
				if counter%2 != 0 {
					temp += "@ "
				} else if counter%2 == 0 {
					temp += "$ "
				}
			}
		}
		temp += "\n"
	}

	return temp
}
func mainr() {
	fmt.Println("Soal Nomor 2")

	fmt.Print(drawSymbolicColsBox(5))
}
