package _2974_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-number-game

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{5, 4, 2, 3},
			res:  []int{3, 2, 5, 4},
		},
		{
			name: "2",
			nums: []int{2, 5},
			res:  []int{5, 2},
		},
		{
			name: "3",
			nums: []int{5, 14, 22, 7, 33, 12, 4, 3},
			res:  []int{4, 3, 7, 5, 14, 12, 33, 22},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberGame(c.nums))
		})
	}
}

func numberGame(nums []int) []int {
	slices.Sort(nums)
	for i := 0; i < len(nums); i += 2 {
		nums[i], nums[i+1] = nums[i+1], nums[i]
	}

	return nums
}
