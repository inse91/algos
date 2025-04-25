package _2089_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-target-indices-after-sorting-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		nums   []int
		target int
		res    []int
	}{
		{
			name:   "1",
			nums:   []int{1, 2, 5, 2, 3},
			target: 2,
			res:    []int{1, 2},
		},
		{
			name:   "2",
			nums:   []int{1, 2, 5, 2, 3},
			target: 3,
			res:    []int{3},
		},
		{
			name:   "3",
			nums:   []int{1, 2, 5, 2, 3},
			target: 5,
			res:    []int{4},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, targetIndices(c.nums, c.target))
		})
	}
}

func targetIndices(nums []int, target int) []int {
	var (
		less int
		eq   int
	)
	for _, num := range nums {
		if num < target {
			less++
		}

		if num == target {
			eq++
		}
	}

	res := make([]int, 0, eq)
	for range eq {
		res = append(res, less)
		less++
	}

	return res
}
