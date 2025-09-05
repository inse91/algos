package _2824_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-pairs-whose-sum-is-less-than-target

func Test(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		res    int
	}{
		{
			name:   "1",
			nums:   []int{-1, 1, 2, 3, 1},
			target: 2,
			res:    3,
		},
		{
			name:   "2",
			nums:   []int{-6, 2, 5, -2, -7, -1, 3},
			target: -2,
			res:    10,
		},
		{
			name:   "3",
			nums:   []int{-1, 1, 2, 3, 1, -4, 15, 3, 9, -4},
			target: -2,
			res:    16,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countPairs(c.nums, c.target))
		})
	}
}

func countPairs(nums []int, target int) int {
	slices.Sort(nums)

	s := make(map[[2]int]struct{})
	var c int
	for i, num1 := range nums {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			num2 := nums[j]
			if num1+num2 >= target {
				break
			}

			if _, ok := s[[2]int{i, j}]; ok {
				continue
			}

			if _, ok := s[[2]int{i, j}]; ok {
				continue
			}

			c++
			s[[2]int{j, i}] = struct{}{}
		}
	}

	return c
}
