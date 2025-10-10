package _892_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/surface-area-of-3d-shapes

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
			res: 34,
		},
		{
			name: "2",
			grid: [][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			res: 32,
		},
		{
			name: "3",
			grid: [][]int{
				{2, 2, 2},
				{2, 1, 2},
				{2, 2, 2},
			},
			res: 46,
		},
		{
			name: "24",
			grid: [][]int{
				{1, 0},
				{5, 4},
			},
			res: 36,
		},
		{
			name: "25",
			grid: [][]int{
				{1, 0},
				{4, 5},
			},
			res: 36,
		},
		{
			name: "26",
			grid: [][]int{
				{1, 0},
				{4, 6},
			},
			res: 40,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, surfaceArea(c.grid))
		})
	}
}

func surfaceArea(grid [][]int) int {
	var srf int
	for i, row := range grid {
		for j, current := range row {
			if current == 0 {
				continue
			}

			var (
				top  int
				down int
			)
			if i != 0 {
				top = grid[i-1][j]
			}
			if i != len(row)-1 {
				down = grid[i+1][j]
			}

			var (
				left  int
				right int
			)
			if j != 0 {
				left = grid[i][j-1]
			}
			if j != len(row)-1 {
				right = grid[i][j+1]
			}

			curSrf := 2 // top + down of self
			if current > top {
				curSrf += current - top
			}
			if current > down {
				curSrf += current - down
			}
			if current > right {
				curSrf += current - right
			}
			if current > left {
				curSrf += current - left
			}

			srf += curSrf
		}
	}

	return srf
}
