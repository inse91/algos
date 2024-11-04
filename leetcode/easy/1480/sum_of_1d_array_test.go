package _1480_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/shuffle-the-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 4},
			res:  []int{1, 3, 6, 10},
		},
		{
			name: "2",
			nums: []int{1, 1, 1, 1, 1},
			res:  []int{1, 2, 3, 4, 5},
		},
		{
			name: "3",
			nums: []int{3, 1, 2, 10, 1},
			res:  []int{3, 4, 6, 16, 17},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, runningSum(c.nums))
		})
	}

}

func runningSum(nums []int) []int {
	var acc int
	for i := range nums {
		acc += nums[i]
		nums[i] = acc
	}

	return nums
}
