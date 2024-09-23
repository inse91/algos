package _1351_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtrx [][]int
		res  int
	}{
		{
			name: "1",
			mtrx: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			res:  8,
		},
		{
			name: "2",
			mtrx: [][]int{{3, 2}, {1, 0}},
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countNegatives(c.mtrx))
		})
	}
}

func countNegatives(grid [][]int) int {
	var res int
	for _, ints := range grid {
		for i := range ints {
			if ints[i] < 0 {
				res += len(ints) - i
				break
			}
		}
	}

	return res
}
