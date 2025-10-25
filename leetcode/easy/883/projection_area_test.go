package _883_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/projection-area-of-3d-shapes

func Test(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		res  int
	}{
		{
			name: "1",
			grid: [][]int{
				{1, 2},
				{3, 4},
			},
			res: 17,
		},
		{
			name: "2",
			grid: [][]int{
				{2},
			},
			res: 5,
		},
		{
			name: "3",
			grid: [][]int{
				{1, 0},
				{0, 2},
			},
			res: 8,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, projectionArea(c.grid))
		})
	}
}

func projectionArea(grid [][]int) int {
	var area int
	for i, row := range grid {
		var (
			mx int // x-axis
			my int // y-axis
		)
		for j, num := range row {
			mx = max(mx, num)
			my = max(my, grid[j][i])

			if num != 0 {
				area += 1 // z-axis
			}
		}

		area += mx
		area += my
	}

	return area
}
