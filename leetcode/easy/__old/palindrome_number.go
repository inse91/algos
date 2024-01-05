package __old

import "fmt"

func isPalNew(x int) bool {
	if x < 0 {
		return false
	}
	var revX int
	copyX := x
	for copyX != 0 {
		revX = revX*10 + copyX%10
		copyX /= 10
	}

	return x == revX
}

func isPalindrome(x int) bool {
	rev := make([]int, 0)
	num := x

	for num > 0 {
		digit := num % 10
		rev = append(rev, digit)
		num /= 10
	}

	num = 0
	for mul, dec := 1, len(rev)-1; dec >= 0; dec-- {
		num += rev[dec] * mul
		mul *= 10
	}

	if num == x {
		return true
	}
	return false
}

func RunPalNum() {

	fmt.Println(isPalNew(123))
}
