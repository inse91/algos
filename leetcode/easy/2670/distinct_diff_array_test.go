package _2670_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-distinct-difference-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 4, 5},
			res:  []int{-3, -1, 1, 3, 5},
		},
		{
			name: "2",
			nums: []int{3, 2, 3, 4, 2},
			res:  []int{-2, -1, 0, 2, 3},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, distinctDifferenceArray(c.nums))
		})
	}
}

func distinctDifferenceArray(nums []int) []int {
	left := make(map[int]struct{}, len(nums))
	right := make(map[int]int, len(nums))
	for _, num := range nums {
		right[num]++
	}

	for i, num := range nums {
		// add to left
		left[num] = struct{}{}

		// dec in tight
		right[num]--
		if right[num] == 0 {
			delete(right, num)
		}

		nums[i] = len(left) - len(right)
	}

	return nums
}
