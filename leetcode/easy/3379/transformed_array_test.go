package _3379_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/transformed-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{3, -2, 1, 1},
			res:  []int{1, 1, 1, 3},
		},
		{
			name: "2",
			nums: []int{-1, 4, -1},
			res:  []int{-1, -1, 4},
		},
		{
			name: "22",
			nums: []int{-7, 10, -1},
			res:  []int{-1, -1, 10},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, constructTransformedArray(c.nums))
		})
	}
}

func constructTransformedArray(nums []int) []int {
	res := make([]int, len(nums))
	for i, num := range nums {
		if num == 0 {
			continue
		}

		newIdx := (i + num) % len(nums)
		if newIdx < 0 {
			newIdx = len(nums) + newIdx
		}

		newIdx = newIdx % len(nums)
		res[i] = nums[newIdx]
	}

	return res
}
