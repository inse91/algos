package _3512_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			k:    5,
			nums: []int{3, 9, 7},
			res:  4,
		},
		{
			name: "2",
			k:    4,
			nums: []int{4, 1, 3},
			res:  0,
		},
		{
			name: "3",
			k:    6,
			nums: []int{3, 2},
			res:  5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minOperations(c.nums, c.k))
		})
	}
}

func minOperations(nums []int, k int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}

	return sum % k
}
