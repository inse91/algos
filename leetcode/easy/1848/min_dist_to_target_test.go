package _1848_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-distance-to-the-target-element

func Test(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		start  int
		res    int
	}{
		{
			name:   "1",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			start:  3,
			res:    1,
		},
		{
			name:   "2",
			nums:   []int{1},
			target: 1,
			start:  0,
			res:    0,
		},
		{
			name:   "3",
			nums:   []int{1, 1, 1, 1, 1, 1},
			start:  0,
			target: 1,
			res:    0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, getMinDistance(c.nums, c.target, c.start))
		})
	}
}

func getMinDistance(nums []int, target int, start int) int {
	if nums[start] == target {
		return 0
	}

	dist := 1
	for {
		leftIdx := start - dist
		rightIdx := start + dist

		if leftIdx >= 0 && nums[leftIdx] == target {
			return dist
		}

		if rightIdx < len(nums) && nums[rightIdx] == target {
			return dist
		}

		dist++
	}
}
