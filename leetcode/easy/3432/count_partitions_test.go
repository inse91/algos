package _3432_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-partitions-with-even-sum-difference24zducyh

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{10, 10, 3, 7, 6},
			res:  4,
		},
		{
			name: "2",
			nums: []int{1, 2, 2},
			res:  0,
		},
		{
			name: "3",
			nums: []int{2, 4, 6, 8},
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countPartitions(c.nums))
		})
	}
}

func countPartitions(nums []int) int {
	var right int
	for _, num := range nums {
		right += num
	}

	var left, res int
	for _, num := range nums[:len(nums)-1] {
		left += num
		right -= num

		if (left-right)%2 == 0 {
			res++
		}
	}

	return res
}
