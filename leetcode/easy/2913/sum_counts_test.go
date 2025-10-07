package _2913_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 1},
			res:  15,
		},
		{
			name: "2",
			nums: []int{1, 1},
			res:  3,
		},
		{
			name: "3",
			nums: []int{1},
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumCounts(c.nums))
		})
	}
}

func sumCounts(nums []int) int {
	res := len(nums)
	for start, num := range nums {
		set := map[uint8]struct{}{
			uint8(num): {},
		}
		for end := start + 1; end < len(nums); end++ {
			set[uint8(nums[end])] = struct{}{}
			res += len(set) * len(set)
		}
	}

	return res
}
