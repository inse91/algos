package easy

import "fmt"

func mySqrt(x int) int {

	//res := 0
	//for res*res <= x {
	//	res++
	//}
	//return res - 1

	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r

}

func Run67() {
	input := 50
	sqrt := mySqrt(input)
	fmt.Printf("res: %d", sqrt)
}
