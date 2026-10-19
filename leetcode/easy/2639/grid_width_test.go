package _2639_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-width-of-columns-of-a-grid

func Test(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		res  []int
	}{
		{
			name: "1",
			grid: [][]int{
				{1},
				{22},
				{333},
			},
			res: []int{3},
		},
		{
			name: "2",
			grid: [][]int{
				{-15, 1, 3},
				{15, 7, 12},
				{5, 6, -2},
			},
			res: []int{3, 1, 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, findColumnWidth(c.grid))
		})
	}
}

func findColumnWidth(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		var c int
		for _, col := range grid {
			c = max(c, ln(col[i]))
		}

		res[i] = c
	}

	return res
}

func ln(num int) int {
	return len(strconv.Itoa(num))
}
