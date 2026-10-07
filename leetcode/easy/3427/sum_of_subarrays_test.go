package _3427_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-variable-length-subarrays

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 3, 1},
			res:  11,
		},
		{
			name: "2",
			nums: []int{3, 1, 1, 2},
			res:  13,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, subarraySum(c.nums))
		})
	}
}

func subarraySum(nums []int) int {
	var res int
	for i, num := range nums {
		start := max(0, i-num)
		var inc int
		for j := start; j <= i; j++ {
			inc += nums[j]
		}

		res += inc
	}

	return res
}
