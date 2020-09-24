package main

import "fmt"

func main() {
	// komentar
	fmt.Println("Hello World")
	var firstName string = "John"

	var lastName string
	lastName = "Wick"

	fmt.Println("halo ", firstName, lastName+"!\n")

	// variable tanpa data type
	var notpe = "Fikri"
	lastNf := "Wikl"

	fmt.Println(notpe, lastNf)

	/* Underscore (_) adalah reserved variable yang bisa dimanfaatkan untuk menampung nilai
	yang tidak dipakai. Bisa dibilang variabel ini merupakan keranjang sampah.

	*/

	_ = "belajar golang "

	name := new(string)
	fmt.Println(name)

	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	const message string = "Halo"
	fmt.Printf("message: %s \n", message)

	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)

	fmt.Printf("nilai %d (%t) \n", value, isEqual)

	var point = 8

	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point > 5 {
		fmt.Println("lulus")
	} else if point == 4 {
		fmt.Println("hampir lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}



	func munculSekalis(num string) {
		var temp =[]string{}
		var once, twos int = 0
		var common
		a := strings.Split(num, "")
		sort.Strings(a)
		numSort := strings.Join(a, "")
	  
		for i := 0; i < len(numSort); i++ {
			twos = twos | (once & num[i])
			once = once ^ num[i]
			common = ~(once & twos)
			
		}
		fmt.Println(temp)
	  }


	  int ones = 0, twos = 0; 
        int common_bit_mask; 
  
        for (int i = 0; i < n; i++) { 
            twos = twos | (ones & arr[i]); 
  
            ones = ones ^ arr[i]; 
  
            be removed from 'ones' and 'twos'*/
            common_bit_mask = ~(ones & twos); 
  
            ones &= common_bit_mask; 
  
            twos &= common_bit_mask; 
        } 
        return ones; 
}
