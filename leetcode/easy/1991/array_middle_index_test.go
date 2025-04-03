package _1991_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-middle-index-in-array/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 3, -1, 8, 4},
			res:  3,
		},
		{
			name: "2",
			nums: []int{1, -1, 4},
			res:  2,
		},
		{
			name: "3",
			nums: []int{2, 5},
			res:  -1,
		},
		{
			name: "4",
			nums: []int{1, -2, 2},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findMiddleIndex(c.nums))
		})
	}
}

func findMiddleIndex(nums []int) int {
	var (
		left  int
		right int
	)
	for i := range nums {
		right += nums[i]
	}

	for i, num := range nums {
		right -= num
		if i != 0 {
			left += nums[i-1]
		}

		if left == right {
			return i
		}
	}

	return -1
}
