package _2784_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-array-is-good

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{2, 1, 3},
			res:  false,
		},
		{
			name: "2",
			nums: []int{1, 3, 3, 2},
			res:  true,
		},
		{
			name: "3",
			nums: []int{1, 1},
			res:  true,
		},
		{
			name: "4",
			nums: []int{3, 4, 4, 1, 2, 1},
			res:  false,
		},
		{
			name: "15",
			nums: []int{1, 2, 3, 4, 5, 5},
			res:  true,
		},
		{
			name: "19",
			nums: []int{1, 2, 3, 7, 5, 5},
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isGood(c.nums))
		})
	}
}

func isGood(nums []int) bool {
	if len(nums) == 1 {
		return false
	}

	slices.Sort(nums)
	maxElemShouldBe := len(nums) - 1
	last := nums[len(nums)-1]
	secondToLast := nums[len(nums)-2]
	if last != secondToLast || secondToLast != maxElemShouldBe {
		return false
	}

	firstPart := nums[:len(nums)-2]
	for i, num := range firstPart {
		if num != i+1 {
			return false
		}
	}

	return true
}
