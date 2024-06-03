package _977_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{-4, -1, 0, 3, 10},
			res:  []int{0, 1, 9, 16, 100},
		},
		{
			name: "2",
			nums: []int{-7, -3, 2, 3, 11},
			res:  []int{4, 9, 9, 49, 121},
		},
		{
			name: "3",
			nums: []int{1, 2, 7},
			res:  []int{1, 4, 49},
		},
		{
			name: "4",
			nums: []int{-9},
			res:  []int{81},
		},
		{
			name: "11",
			nums: []int{-5, -3, -2, -1},
			res:  []int{1, 4, 9, 25},
		},
		{
			name: "11_1",
			nums: []int{-5, -3, -2, -1, 0},
			res:  []int{0, 1, 4, 9, 25},
		},
		{
			name: "11_2",
			nums: []int{-5, -3, -2, -1, 2},
			res:  []int{1, 4, 4, 9, 25},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, sortedSquares(c.nums))
		})
	}
}

func sortedSquares(nums []int) []int {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return []int{nums[0] * nums[0]}
	}

	if nums[0] >= 0 {
		for i := range nums {
			nums[i] = pw2(nums[i])
		}

		return nums
	}

	res := make([]int, 0, len(nums))

	if nums[len(nums)-1] < 0 {
		for i := len(nums) - 1; i >= 0; i-- {
			res = append(res, pw2(nums[i]))
		}

		return res
	}

	// find pivot
	pvt := pivot(nums)

	var l, r int
	for {
		if len(res) == len(nums) {
			break
		}

		li, ri := pvt-l-1, pvt+r
		if li < 0 && ri >= len(nums) {
			break
		}

		if li < 0 {
			res = append(res, pw2(nums[ri]))
			r++
			continue
		}

		if ri >= len(nums) {
			res = append(res, pw2(nums[li]))
			l++
			continue
		}

		a, b := abs(nums[li]), abs(nums[ri])

		if a < b {
			res = append(res, pw2(a))
			l++
		} else {
			res = append(res, pw2(b))
			r++
		}
	}

	return res
}

func pivot(nums []int) int {
	for i := range nums {
		if nums[i] >= 0 {
			return i
		}
	}

	return -1
}

func pw2(num int) int {
	return num * num
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
