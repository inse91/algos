package _3151_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/special-array-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1},
			res:  true,
		},
		{
			name: "2",
			nums: []int{2, 1, 4},
			res:  true,
		},
		{
			name: "3",
			nums: []int{4, 3, 1, 6},
			res:  false,
		},
		{
			name: "4",
			nums: []int{2, 4},
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isArraySpecial(c.nums))
		})
	}
}

func isArraySpecial(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	isOdd := nums[0]&1 == 0
	for _, n := range nums[1:] {
		currentIsOdd := n&1 == 0
		if isOdd == currentIsOdd {
			return false
		}

		isOdd = !isOdd
	}

	return true
}
