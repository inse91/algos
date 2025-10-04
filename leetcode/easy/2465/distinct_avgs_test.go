package _2465_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-distinct-averages

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{4, 1, 4, 0, 3, 5},
			res:  2,
		},
		{
			name: "2",
			nums: []int{1, 100},
			res:  1,
		},
		{
			name: "3",
			nums: []int{4, 1, 4, 0, 3, 5, 8, 15},
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, distinctAverages(c.nums))
		})
	}
}

func distinctAverages(nums []int) int {
	slices.Sort(nums)
	m := map[float64]struct{}{}
	for i := 0; i < len(nums)/2; i++ {
		endIdx := len(nums) - i - 1

		start := nums[i]
		end := nums[endIdx]
		m[float64(start+end)/2] = struct{}{}
	}

	return len(m)
}
