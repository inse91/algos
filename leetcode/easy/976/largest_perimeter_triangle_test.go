package _976_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/largest-perimeter-triangle/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 1, 2},
			res:  5,
		},
		{
			name: "2",
			nums: []int{1, 2, 1, 10},
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, largestPerimeter(c.nums))
		})
	}
}

func largestPerimeter(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	slices.Sort(nums)
	for i := len(nums) - 1; i > 1; i-- {
		c := nums[i]
		a, b := nums[i-1], nums[i-2]
		if a+b > c {
			return a + b + c
		}
	}

	return 0
}
