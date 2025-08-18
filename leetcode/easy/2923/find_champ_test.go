package _2923_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-champion-i

func Test(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		res  int
	}{
		{
			name: "1",
			grid: [][]int{
				{0, 1},
				{0, 0},
			},
			res: 0,
		},
		{
			name: "2",
			grid: [][]int{
				{0, 0, 1},
				{1, 0, 1},
				{0, 0, 0},
			},
			res: 1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findChampion(c.grid))
		})
	}
}

func findChampion(grid [][]int) int {
	var (
		mx  int
		idx int
	)

	for i, row := range grid {
		var ones int
		for _, v := range row {
			ones += v
		}

		if ones == len(row)-1 {
			return i
		}

		if ones > mx {
			mx = ones
			idx = i
		}
	}

	return idx
}
