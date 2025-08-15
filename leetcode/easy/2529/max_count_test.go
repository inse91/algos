package _2529_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{-2, -1, -1, 1, 2, 3},
			res:  3,
		},
		{
			name: "2",
			nums: []int{-3, -2, -1, 0, 0, 1, 2},
			res:  3,
		},
		{
			name: "3",
			nums: []int{5, 20, 66, 1314},
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumCount(c.nums))
		})
	}
}

func maximumCount(nums []int) int {
	var pos, neg int
	for _, num := range nums {
		if num == 0 {
			continue
		}
		if num < 0 {
			neg++
			continue
		}

		pos++
	}

	pos = max(pos, neg)

	return pos
}
