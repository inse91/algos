package _2057_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/smallest-index-with-equal-value/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{0, 1, 2},
			res:  0,
		},
		{
			name: "2",
			nums: []int{4, 3, 2, 1},
			res:  2,
		},
		{
			name: "3",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			res:  -1,
		},
		{
			name: "439",
			nums: []int{7, 8, 3, 5, 2, 6, 3, 1, 1, 4, 5, 4, 8, 7, 2, 0, 9, 9, 0, 5, 7, 1, 6},
			res:  21,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, smallestEqual(c.nums))
		})
	}
}

func smallestEqual(nums []int) int {
	for i, num := range nums {
		if i%10 == num {
			return i
		}
	}

	return -1
}
