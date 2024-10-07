package _1389_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		nums  []int
		index []int
		res   []int
	}{
		{
			name:  "1",
			nums:  []int{0, 1, 2, 3, 4},
			index: []int{0, 1, 2, 2, 1},
			res:   []int{0, 4, 1, 3, 2},
		},
		{
			name:  "2",
			nums:  []int{1, 2, 3, 4, 0},
			index: []int{0, 1, 2, 3, 0},
			res:   []int{0, 1, 2, 3, 4},
		},
		{
			name:  "3",
			nums:  []int{1},
			index: []int{0},
			res:   []int{1},
		},
		{
			name:  "9",
			nums:  []int{4, 2, 4, 3, 2},
			index: []int{0, 0, 1, 3, 1},
			res:   []int{2, 2, 4, 4, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//assert.ElementsMatch(t, c.res, createTargetArray(c.nums, c.index))
			assert.Equal(t, c.res, createTargetArray(c.nums, c.index))
		})
	}
}

func createTargetArray(nums []int, index []int) []int {
	res := make([]int, 0, len(nums))
	for i := range index {
		idx := index[i]
		num := nums[i]

		res = slices.Insert(res, idx, num)
	}

	return res
}
