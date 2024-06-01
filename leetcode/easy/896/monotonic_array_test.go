package _896_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/monotonic-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 3},
			res:  true,
		},
		{
			name: "2",
			nums: []int{6, 5, 4, 4},
			res:  true,
		},
		{
			name: "3",
			nums: []int{1, 3, 2},
			res:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isMonotonic(c.nums))
		})
	}
}

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}

	var dec int
	for i := 1; i < len(nums); i++ {
		switch dec {
		case 0:
			switch {
			case nums[i] > nums[i-1]:
				dec = 1
			case nums[i] < nums[i-1]:
				dec = -1
			}
		case -1:
			if nums[i] > nums[i-1] {
				return false
			}
		case 1:
			if nums[i] < nums[i-1] {
				return false
			}
		}
	}

	return true
}
