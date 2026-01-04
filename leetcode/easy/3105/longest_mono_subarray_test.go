package _3105_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-maximum-achievable-number

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 4, 3, 3, 2},
			res:  2,
		},
		{
			name: "2",
			nums: []int{3, 3, 3, 3},
			res:  1,
		},
		{
			name: "3",
			nums: []int{3, 2, 1},
			res:  3,
		},
		{
			name: "19",
			nums: []int{24, 47, 24, 23, 14, 6, 9, 2, 3, 19},
			res:  5,
		},
		{
			name: "27",
			nums: []int{21, 32, 1, 39, 40, 24, 6, 21, 24, 33, 25, 33, 6, 50, 34, 14, 24, 12, 2, 1},
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, longestMonotonicSubarray(c.nums))
		})
	}
}

func longestMonotonicSubarray(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var (
		res  int
		c    = 2
		incr = nums[1] > nums[0]
	)
	if nums[0] == nums[1] {
		c = 1
	}

	for i, num := range nums[1:] {
		if i == 0 {
			continue
		}

		prev := nums[i] // start gap
		if num == prev {
			res = max(res, c)
			c = 1
			continue
		}

		curIncr := num > prev
		if incr && curIncr || !incr && !curIncr {
			c++
			continue
		}

		res = max(res, c)
		c = 2
		incr = curIncr
	}

	return max(res, c)
}
