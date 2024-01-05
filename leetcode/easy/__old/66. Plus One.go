package __old

import (
	"fmt"
)

func plusOne(digits []int) []int {

	if digits == nil || len(digits) == 0 {
		return nil
	}
	lastIdx := len(digits) - 1

	for lastIdx >= 0 && digits[lastIdx] == 9 {
		digits[lastIdx] = 0
		lastIdx--
	}
	if lastIdx == -1 {
		digits = append([]int{1}, digits...)
		return digits
	}
	digits[lastIdx] += 1
	return digits

}

func Run66() {
	input := []int{1, 2, 1, 9}
	//input := []int{9}
	//input := []int{7, 2, 8, 5, 0, 9, 1, 2, 9, 5, 3, 6, 6, 7, 3, 2, 8, 4, 3, 7, 9, 5, 7, 7, 4, 7, 4, 9, 4, 7, 0, 1, 1, 1, 7, 4, 0, 0, 6}
	op := plusOne(input)

	fmt.Printf("res: %#v", op)
}
