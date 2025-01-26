package _1608_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{3, 5},
			res:  2,
		},
		{
			name: "2",
			nums: []int{0, 0},
			res:  -1,
		},
		{
			name: "3",
			nums: []int{0, 4, 3, 0, 4},
			res:  3,
		},
		{
			name: "57",
			nums: []int{1, 0, 0, 6, 4, 9},
			res:  3,
		},
		{
			name: "70",
			nums: []int{3, 6, 7, 7, 0},
			res:  -1,
		},
		{
			name: "73",
			nums: []int{1, 2, 3, 4},
			res:  -1,
		},
		{
			name: "75",
			nums: []int{1, 3, 3, 4},
			res:  3,
		},
		{
			name: "79",
			nums: []int{4, 4, 4, 4},
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, specialArray(c.nums))
		})
	}
}

func specialArray(nums []int) int {
	slices.Sort(nums)
	if nums[0] >= len(nums) {
		return len(nums)
	}

	for num := 1; num < len(nums); num++ {
		if len(nums)-num-1 < 0 {
			continue
		}

		cur := nums[len(nums)-num]
		prev := nums[len(nums)-num-1]

		if cur == prev {
			continue
		}

		if cur >= num && prev < num {
			return num
		}
	}

	return -1
}
