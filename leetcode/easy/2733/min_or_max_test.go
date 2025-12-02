package _2678_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/neither-minimum-nor-maximum

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{3, 2, 1, 4},
			res:  []int{2, 3},
		},
		{
			name: "2",
			nums: []int{1, 2},
			res:  nil,
		},
		{
			name: "3",
			nums: []int{2, 1, 3},
			res:  []int{2},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := findNonMinOrMax(c.nums)
			if res == -1 {
				assert.Empty(t, c.res)
				return
			}
			assert.Contains(t, c.res, res)
		})
	}
}

func findNonMinOrMax(nums []int) int {
	if len(nums) <= 2 {
		return -1
	}

	var (
		mx = -1
		mn = 101
	)

	for _, num := range nums {
		mx = max(mx, num)
		mn = min(mn, num)

	}
	for _, num := range nums {
		if num < mx && num > mn {
			return num
		}
	}

	return -1
}
