package _2460_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/apply-operations-to-an-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 1, 1, 0},
			res:  []int{1, 4, 2, 0, 0, 0},
		},
		{
			name: "2",
			nums: []int{0, 1},
			res:  []int{1, 0},
		},
		{
			name: "3",
			nums: []int{1, 4, 0, 13, 2, 2, 23, 0, 1, 1, 0},
			res:  []int{1, 4, 13, 4, 23, 2, 0, 0, 0, 0, 0},
		},
		{
			name: "4",
			nums: []int{1, 4, 0, 13, 2, 2, 23, 0, 1, 1, 0, 4, 4, 14},
			res:  []int{1, 4, 13, 4, 23, 2, 8, 14, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, applyOperations(c.nums))
		})
	}
}

func applyOperations(nums []int) []int {
	res := make([]int, 0, len(nums))
	var zeroCount int
	for i := 0; i <= len(nums)-1; i++ {
		cur := nums[i]
		if nums[i] == 0 { // num is zero
			zeroCount++
			continue
		}

		if i == len(nums)-1 { // last non-zero num goes to the result anyway
			res = append(res, cur)
			continue
		}

		next := nums[i+1]
		if cur == next {
			cur *= 2
			nums[i+1] = 0
		}

		res = append(res, cur)
	}

	return append(res, make([]int, zeroCount)...)
}
