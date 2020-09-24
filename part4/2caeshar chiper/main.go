package main

import "fmt"

// Rotate Latin letters by the shift amount.
func caesarChiper(text string, shift int) string {
	shift = (shift%26 + 26) % 26 // [0, 25]
	b := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		t := text[i]
		var a int
		switch {
		case 'a' <= t && t <= 'z':
			a = 'a'
		case 'A' <= t && t <= 'Z':
			a = 'A'
		default:
			b[i] = t
			continue
		}
		b[i] = byte(a + ((int(t)-a)+shift)%26)
	}
	return string(b)
}

// Encode using Caesar Cipher.
func Encode(plain string, shift int) (cipher string) {
	return caesarChiper(plain, shift)
}

// Decode using Caesar Cipher.
func Decode(cipher string, shift int) (plain string) {
	return caesarChiper(cipher, -shift)
}

func main() {
	a := caesarChiper("abcdefghijklmnopqrstufwxyz", 1000)
	fmt.Println(a)
}

// func main() {
// 	a := caesarChiper("abcdefghijklmnopqrstufwxyz", 1000)
// 	fmt.Println(a)
// 	b := Decode("zab", -1)
// 	fmt.Println(b)
// }
