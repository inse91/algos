package _2016_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-difference-between-increasing-elements

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{7, 1, 5, 4},
			res:  4,
		},
		{
			name: "2",
			nums: []int{9, 4, 3, 2},
			res:  -1,
		},
		{
			name: "3",
			nums: []int{1, 5, 2, 10},
			res:  9,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumDifference(c.nums))
		})
	}
}

func maximumDifference(nums []int) int {
	res := -1
	for i, num := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[j] <= num {
				continue
			}

			res = max(res, nums[j]-num)
		}
	}

	return res
}
