package _1252_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/cells-with-odd-values-in-a-matrix/description/

func TestName(t *testing.T) {
	cases := []struct {
		name string
		m    int
		n    int
		mtrx [][]int
		res  int
	}{
		{
			name: "1",
			m:    2,
			n:    3,
			mtrx: [][]int{{0, 1}, {1, 1}},
			res:  6,
		},
		{
			name: "2",
			m:    2,
			n:    2,
			mtrx: [][]int{{1, 1}, {0, 0}},
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, oddCells(c.m, c.n, c.mtrx))
		})
	}
}

func oddCells(m int, n int, indices [][]int) int {
	mtrx := make([][]int, m)
	for i := range mtrx {
		mtrx[i] = make([]int, n)
	}

	for _, indc := range indices {
		r := indc[0]
		c := indc[1]

		for i := range mtrx[r] {
			mtrx[r][i]++
		}

		for i := range mtrx {
			mtrx[i][c]++
		}
	}

	var res int
	for _, ints := range mtrx {
		for _, num := range ints {
			res += num % 2
		}
	}

	return res
}
