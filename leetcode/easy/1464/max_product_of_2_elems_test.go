package _1464_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/reformat-date/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{3, 4, 5, 2},
			res:  12,
		},
		{
			name: "2",
			nums: []int{1, 5, 4, 5},
			res:  16,
		},
		{
			name: "3",
			nums: []int{3, 7},
			res:  12,
		},
		{
			name: "60",
			nums: []int{10, 2, 5, 2},
			res:  36,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxProduct(c.nums))
		})
	}
}

func maxProduct(nums []int) int {
	var cur, prev int
	for _, num := range nums {
		if num >= cur {
			cur, prev = num, cur
			continue
		}

		if num > prev {
			prev = num
		}
	}

	return (cur - 1) * (prev - 1)
}
