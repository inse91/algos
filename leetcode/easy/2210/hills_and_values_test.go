package _2210_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-hills-and-valleys-in-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 4, 1, 1, 6, 5},
			res:  3,
		},
		{
			name: "2",
			nums: []int{6, 6, 5, 5, 4, 1},
			res:  0,
		},
		{
			name: "13",
			nums: []int{2, 4, 1, 6, 5},
			res:  3,
		},
		{
			name: "51",
			nums: []int{8, 2, 5, 7, 7, 2, 10, 3, 6, 2},
			res:  6,
		},
		{
			name: "62",
			nums: []int{1, 2, 1, 2, 1, 2, 1},
			res:  5,
		},
		{
			name: "63",
			nums: []int{1, 2, 1, 2, 1},
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countHillValley(c.nums))
		})
	}
}

func countHillValley(nums []int) int {
	var res int
	prev := nums[0]
	for i := 1; i <= len(nums)-2; i++ {
		cur := nums[i]
		if cur == prev {
			continue
		}

		next := nums[i+1]
		if next == cur {
			continue
		}

		if min(cur, prev, next) < min(prev, next) {
			res++
		}

		if max(cur, prev, next) > max(prev, next) {
			res++
		}

		prev = cur
	}

	return res
}
