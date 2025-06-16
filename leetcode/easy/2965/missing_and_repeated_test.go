package _2965_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-missing-and-repeated-values

func Test(t *testing.T) {
	cases := []struct {
		name string
		grid [][]int
		res  []int
	}{
		{
			name: "1",
			grid: [][]int{
				{1, 3},
				{2, 2},
			},
			res: []int{2, 4},
		},
		{
			name: "2",
			grid: [][]int{
				{9, 1, 7},
				{8, 9, 2},
				{3, 4, 6},
			},
			res: []int{9, 5},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findMissingAndRepeatedValues(c.grid))
		})
	}
}

func findMissingAndRepeatedValues(grid [][]int) []int {
	if len(grid) == 0 {
		panic("nil grid")
	}

	var res1 int
	slc := make([]int, len(grid)*len(grid[0]))

	for _, g := range grid {
		for _, i := range g {
			i--
			if slc[i] == 1 {
				res1 = i + 1
				continue
			}

			slc[i] = 1
		}
	}

	for i, c := range slc {
		if c == 0 {
			return []int{res1, i + 1}
		}
	}

	return nil
}
