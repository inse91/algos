package _1051_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/height-checker/description/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		heights []int
		res     int
	}{
		{
			name:    "1",
			heights: []int{1, 1, 4, 2, 1, 3},
			res:     3,
		},
		{
			name:    "2",
			heights: []int{5, 1, 2, 3, 4},
			res:     5,
		},
		{
			name:    "3",
			heights: []int{1, 2, 3, 4, 5},
			res:     0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, heightChecker(c.heights))
		})
	}
}

func heightChecker(heights []int) int {
	heightCopy := slices.Clone(heights)
	slices.Sort(heights)

	var res int
	for i := range heights {
		if heights[i] != heightCopy[i] {
			res++
		}
	}

	return res
}
