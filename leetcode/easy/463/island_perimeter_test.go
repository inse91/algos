package _463_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIslandPerimeter(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		res  int
	}{
		{
			name: "1",
			grid: [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}},
			res:  16,
		},
		{
			name: "2",
			grid: [][]int{{1}},
			res:  4,
		},
		{
			name: "3",
			grid: [][]int{{1, 0}},
			res:  4,
		},
		{
			name: "4",
			grid: [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			res:  12,
		},
		{
			name: "5",
			grid: [][]int{{1, 1}, {1, 1}},
			res:  8,
		},
		{
			name: "6",
			grid: [][]int{{0, 1}},
			res:  4,
		},
		{
			name: "7",
			grid: [][]int{{0}, {1}},
			res:  4,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			//t.Parallel()
			res := islandPerimeter(c.grid)
			assert.Equal(t, c.res, res)
		})
	}
}

func islandPerimeter(grid [][]int) (per int) {
	for i, arr := range grid {
		for j, num := range arr {
			if num == 0 {
				continue
			}
			// left
			if j == 0 || arr[j-1] == 0 {
				per++
			}
			//right
			if j == len(arr)-1 || arr[j+1] == 0 {
				per++
			}
			// up
			if i == 0 || grid[i-1][j] == 0 {
				per++
			}
			// down

			if i == len(grid)-1 || grid[i+1][j] == 0 {
				per++
			}
		}
	}
	return per
}
