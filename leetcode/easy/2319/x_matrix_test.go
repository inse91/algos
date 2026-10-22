package _2319_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-if-matrix-is-x-matrix

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtrx [][]int
		res  bool
	}{
		{
			name: "1",
			mtrx: [][]int{
				{2, 0, 0, 1},
				{0, 3, 1, 0},
				{0, 5, 2, 0},
				{4, 0, 0, 2},
			},
			res: true,
		},
		{
			name: "2",
			mtrx: [][]int{
				{5, 7, 0},
				{0, 3, 1},
				{0, 5, 0},
			},
			res: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, checkXMatrix(c.mtrx))
		})
	}
}

func checkXMatrix(grid [][]int) bool {
	ln := len(grid)
	for i, row := range grid {
		for j, num := range row {
			if isDiagonal(i, j, ln) {
				if num != 0 {
					continue
				}

				return false
			}

			if num != 0 {
				return false
			}
		}
	}

	return true
}

func isDiagonal(i, j, ln int) bool {
	if i == j {
		return true
	}

	if i+j+1 == ln {
		return true
	}

	return false
}
