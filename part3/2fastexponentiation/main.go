package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(new(big.Int).Exp(big.NewInt(5), big.NewInt(300), nil))
}

func powBig(a, n int) *big.Int {
	tmp := big.NewInt(int64(a))
	result := big.NewInt(1)
	for n > 0 {
		temp := new(big.Int)
		if n%2 == 1 {
			temp.Mul(result, tmp)
			result = temp
		}
		temp = new(big.Int)
		temp.Mul(tmp, tmp)
		tmp = temp
		n /= 2
	}
	return result
}
