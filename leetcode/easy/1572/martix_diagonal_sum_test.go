package _1572_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/matrix-diagonal-sum/

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtrx [][]int
		res  int
	}{
		{
			name: "1",
			mtrx: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			res: 25,
		},
		{
			name: "2",
			mtrx: [][]int{
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
				{1, 1, 1, 1},
			},
			res: 8,
		},
		{
			name: "3",
			mtrx: [][]int{
				{5},
			},
			res: 5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, diagonalSum(c.mtrx))
		})
	}
}

func diagonalSum(mat [][]int) int {
	var (
		left  = 0
		right = len(mat) - 1
		sum   int
	)

	for i := range mat {
		if left == right {
			sum += mat[i][left]
			left++
			right--
			continue
		}

		sum += mat[i][left] + mat[i][right]
		left++
		right--
	}

	return sum
}
