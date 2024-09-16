package _1260_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/shift-2d-grid/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		k    int
		res  [][]int
	}{
		{
			name: "1",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    1,
			res:  [][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}},
		},
		{
			name: "2",
			grid: [][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}},
			k:    4,
			res:  [][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}},
		},
		{
			name: "3",
			grid: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			k:    9,
			res:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			for i, ints := range shiftGrid(c.grid, c.k) {
				assert.EqualValues(t, ints, c.res[i])
			}
		})
	}
}

func shiftGrid(grid [][]int, k int) [][]int {
	sl := make([]int, 0, len(grid)*len(grid[0]))
	for _, ints := range grid {
		for _, n := range ints {
			sl = append(sl, n)
		}
	}

	sl = shiftSlice(sl, k)

	var p int
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = sl[p]
			p++
		}
	}

	return grid
}

func shiftSlice(sl []int, k int) []int {
	k = k % len(sl)
	k = len(sl) - k
	if k == 0 {
		return sl
	}

	return append(sl[k:], sl[:k]...)
}
