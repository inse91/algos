package easy

func searchInsert(nums []int, target int) int {

	for i, num := range nums {
		if num >= target {
			return i
		}
	}

	return len(nums)

}

func Run35() {
	println(searchInsert([]int{1, 3, 5, 6}, 2))
}
