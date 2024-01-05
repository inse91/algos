package __old

import "fmt"

func removeElement(nums []int, val int) int {

	i := 0
	for _, num := range nums {
		if num != val {
			nums[i] = num
			i++
		}
	}
	return i
}

func Run27RemoveELem() {
	input := []int{1, 2, 2, 3, 2, 4}
	x := removeElement(input, 2)
	fmt.Printf("result: %d", x)
}
