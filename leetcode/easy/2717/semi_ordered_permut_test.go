package _2717_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/semi-ordered-permutation

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 1, 4, 3},
			res:  2,
		},
		{
			name: "2",
			nums: []int{2, 4, 1, 3},
			res:  3,
		},
		{
			name: "3",
			nums: []int{1, 3, 4, 2, 5},
			res:  0,
		},
		{
			name: "13",
			nums: []int{1, 3},
			res:  0,
		},
		{
			name: "17",
			nums: []int{5, 2},
			res:  1,
		},
		{
			name: "23",
			nums: []int{5, 2, 4, 1, 3},
			res:  6,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, semiOrderedPermutation(c.nums))
		})
	}
}

func semiOrderedPermutation(nums []int) int {
	switch len(nums) {
	case 0, 1:
		return 0
	case 2:
		if nums[0] > nums[1] {
			return 1
		}

		return 0
	default:
	}

	var (
		mn, mx       = 51, 0
		mnIdx, mxIdx = -1, -1
	)

	for i, num := range nums {
		if num < mn {
			mn = num
			mnIdx = i
		}
		if num > mx {
			mx = num
			mxIdx = i
		}
	}

	if mnIdx == 0 && mxIdx == len(nums)-1 {
		return 0
	}

	minToLeft := mnIdx
	maxToRight := len(nums) - mxIdx - 1
	if mnIdx != 0 && mxIdx < mnIdx {
		maxToRight--
	}

	return minToLeft + maxToRight
}
