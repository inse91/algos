package _2869_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-operations-to-collect-elements

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{3, 1, 5, 4, 2},
			k:    2,
			res:  4,
		},
		{
			name: "2",
			nums: []int{3, 1, 5, 4, 2},
			k:    5,
			res:  5,
		},
		{
			name: "3",
			nums: []int{3, 2, 5, 3, 1},
			k:    3,
			res:  4,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minOperations(c.nums, c.k))
		})
	}
}

func minOperations(nums []int, k int) int {
	if len(nums) == k {
		return k
	}

	set := make(map[int]int, k)
	for i, num := range nums {
		if num > k {
			continue
		}

		set[num] = i
	}

	minIdx := len(nums)
	for i := 1; i <= k; i++ {
		minIdx = min(minIdx, set[i])
	}

	return len(nums) - minIdx
}
