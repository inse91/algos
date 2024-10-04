package _1380_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/lucky-numbers-in-a-matrix/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtrx [][]int
		res  []int
	}{
		{
			name: "1",
			mtrx: [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}},
			res:  []int{15},
		},
		{
			name: "2",
			mtrx: [][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}},
			res:  []int{12},
		},
		{
			name: "3",
			mtrx: [][]int{{7, 8}, {1, 2}, {5, 6}},
			res:  []int{7},
		},
		{
			name: "3_1",
			mtrx: [][]int{{7, 8, 9}, {1, 2, 3}},
			res:  []int{7},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, luckyNumbers(c.mtrx))
		})
	}
}

func luckyNumbers(mtrx [][]int) []int {
	minRows := make([]int, 0, len(mtrx))
	maxCols := make([]int, len(mtrx[0]))

	for i := range mtrx {
		row := mtrx[i]
		var mn int = 1e5 + 1
		for j, num := range row {
			mn = min(mn, num)
			maxCols[j] = max(maxCols[j], num)
		}
		minRows = append(minRows, mn)
	}

	var res []int
	for i := range mtrx {
		for j := range mtrx[i] {
			num := mtrx[i][j]

			if num == minRows[i] && num == maxCols[j] {
				res = append(res, num)
			}
		}
	}

	return res
}
