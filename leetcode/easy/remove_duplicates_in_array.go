package easy

import "fmt"

func removeDuplicates(nums []int) (res int) {
	//m := make(map[int]bool)
	//for _, num := range nums {
	//	m[num] = true
	//}
	//nums = make([]int, 0, len(m))
	//for k, _ := range m {
	//	nums = append(nums, k)
	//}
	//fmt.Printf("%+v", nums)
	//return len(m)
	//cur := 101
	//for i := 0; i < len(nums); {
	//
	//	if cur == nums[i] {
	//		nums = append(nums[:i], nums[i+1:]...)
	//		continue
	//	}
	//	cur = nums[i]
	//	i++
	//}
	//return len(nums)

	if len(nums) == 1 {
		return 1
	}
	count := 0
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i] {
			count++
		} else {
			nums[i-count] = nums[i]
		}
	}
	fmt.Printf("%+v", nums[:len(nums)-count])
	return len(nums) - count
}

func RunRemoveDuplicates() {
	input := []int{1, 1, 2, 3, 3, 3}
	fmt.Println(removeDuplicates(input))
}
