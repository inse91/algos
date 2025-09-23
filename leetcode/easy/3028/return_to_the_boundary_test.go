package _3028_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/ant-on-the-boundary

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{2, 3, -5},
			res:  1,
		},
		{
			name: "2",
			nums: []int{3, 2, -3, -4},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, returnToBoundaryCount(c.nums))
		})
	}
}

func returnToBoundaryCount(nums []int) int {
	var (
		c   int
		pos int
	)
	for _, num := range nums {
		pos += num
		if pos == 0 {
			c++
		}
	}

	return c
}
