package _1929_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/concatenation-of-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, 1},
			res:  []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "2",
			nums: []int{1, 3, 2, 1},
			res:  []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, getConcatenation(c.nums))
		})
	}
}

func getConcatenation(nums []int) []int {
	return append(nums, nums...)
}
