package _1752_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{3, 4, 5, 1, 2},
			res:  true,
		},
		{
			name: "2",
			nums: []int{2, 1, 3, 4},
			res:  false,
		},
		{
			name: "2_1",
			nums: []int{2, 1, 1, 2},
			res:  true,
		},
		{
			name: "3",
			nums: []int{1, 2, 3},
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, check(c.nums))
		})
	}
}

func check(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var wasRotated bool
	for i := 1; i < len(nums); i++ {
		cur, prev := nums[i], nums[i-1]
		if cur >= prev {
			continue
		}

		if wasRotated {
			return false
		}

		wasRotated = true
	}

	return !wasRotated || nums[len(nums)-1] <= nums[0]
}
