package main

import (
	"fmt"
	"strings"
)

func main() {
	pyramid(5)
}

func pyramid(number int) {
	row := "* "
	blankSpace := number - 1

	for i := 0; i < 5; i++ {
		str := strings.Repeat(" ", blankSpace) + row
		fmt.Println(str)
		if blankSpace-1 != -1 {
			blankSpace--
		}
		row += "* "
	}
}
