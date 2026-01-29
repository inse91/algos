package _2335_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-amount-of-time-to-fill-cups

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 4, 2},
			res:  4,
		},
		{
			name: "12",
			nums: []int{5, 4, 1},
			res:  5,
		},
		{
			name: "3",
			nums: []int{5, 0, 0},
			res:  5,
		},
		{
			name: "27",
			nums: []int{0, 0, 0},
			res:  0,
		},
		{
			name: "33",
			nums: []int{7, 12, 1},
			res:  12,
		},
		{
			name: "36",
			nums: []int{7, 12, 11},
			res:  15,
		},
		{
			name: "23",
			nums: []int{4, 4, 4},
			res:  6,
		},
		{
			name: "22",
			nums: []int{5, 4, 2},
			res:  6,
		},
		{
			name: "245",
			nums: []int{3, 4, 4},
			res:  6,
		},
		{
			name: "250",
			nums: []int{8, 6, 5},
			res:  10,
		},
		{
			name: "2",
			nums: []int{5, 4, 4},
			res:  7,
		},
		{
			name: "279",
			nums: []int{97, 84, 84},
			res:  133,
		},
		{
			name: "263",
			nums: []int{7, 4, 4},
			res:  8,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := fillCups(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func fillCups(amount []int) int {
	slices.Sort(amount)
	var (
		mx = amount[2]
		h  = amount[1]
		l  = amount[0]
	)

	// max is too large
	if mx >= h+l {
		return mx
	}

	mx -= l
	sm := mx + h + 1
	return l + sm/2
}
