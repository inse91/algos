package __old

import (
	"fmt"
	"math/big"
)

func addBinary(a string, b string) string {

	num1, num2, summ := new(big.Int), new(big.Int), new(big.Int)

	num1.SetString(a, 2)
	num2.SetString(b, 2)

	summ.Add(num1, num2)
	s := summ.Text(2)

	return s
}

func RunAddBinary() {
	a := "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101"
	b := "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"
	fmt.Println(addBinary(a, b))
}
