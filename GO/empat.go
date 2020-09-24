package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

func ubahHuruf(text string) string {

	var words string
	var temp string = ""
	words = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var convertedText = strings.ToUpper(text)

	for _, index := range convertedText {
		newIndexOfText := strings.IndexAny(words, string(index)) + 10

		if newIndexOfText > 26 {
			newIndexOfText %= 26
		}

		if string(index) == " " {
			temp += " "
		} else {
			temp += string(words[newIndexOfText])
		}
	}

	return temp
}
func maisn() {
	fmt.Println("Hello World")
	fmt.Println(ubahHuruf("SEPULSa OKE"))
	// munculSekali("")
	// munculSekali("-1")
	munculSekalis("1122")
	munculSekali("76523752")
}

func munculSekalis(num string) {
	var temp = []string{}

	a := strings.Split(num, "")
	sort.Strings(a)
	numSort := strings.Join(a, "")

	for i := 0; i < len(numSort); i++ {
		if i == 0 && numSort[i] != numSort[i+1] {
			temp = append(temp, string(numSort[i]))
		} else if i == 0 && numSort[i] == numSort[i+1] {
			temp = []string{}
		} else if i == len(numSort)-1 && numSort[i] != numSort[i-1] {
			temp = append(temp, string(numSort[i]))
		} else if numSort[i] != numSort[i-1] && numSort[i] != numSort[i+1] {
			temp = append(temp, string(numSort[i]))
		}
	}
	fmt.Println(temp)
}

func munculSekali(num string) {
	var (
		// result = ""
		result = []string{}
	)

	if len(num) <= 0 {
		fmt.Println("Invalid Input")
		return
	}

	for _, checkString := range num {
		if !unicode.IsNumber(checkString) {
			fmt.Println("Invalid Input")
			return
		}
	}

	a := strings.Split(num, "")
	sort.Strings(a)
	numSort := strings.Join(a, "")

	// if unicode.IsLetter(numSort) {
	//   return
	// }

	for i := 0; i < len(numSort); i++ {
		if i == 0 && numSort[i] != numSort[i+1] {
			// result += string(numSort[i])
			result = append(result, string(numSort[i]))
		} else if i == 0 && numSort[i] == numSort[i+1] {
			// result += string(numSort[i])
			result = []string{}
		} else if i == len(numSort)-1 && numSort[i] != numSort[i-1] {
			// result += string(numSort[i])
			result = append(result, string(numSort[i]))
		} else if numSort[i] != numSort[i-1] && numSort[i] != numSort[i+1] {
			// result += string(numSort[i])
			result = append(result, string(numSort[i]))
		}
	}
	fmt.Println(result)
}
