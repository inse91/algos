package _566_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReshapeMatrix(t *testing.T) {
	cases := []struct {
		name string
		mat  [][]int
		r    int
		c    int
		res  [][]int
	}{
		{
			name: "1",
			mat:  [][]int{{1, 2}, {3, 4}},
			r:    1,
			c:    4,
			res:  [][]int{{1, 2, 3, 4}},
		},
		{
			name: "2",
			mat:  [][]int{{1, 2}, {3, 4}},
			r:    2,
			c:    2,
			res:  [][]int{{1, 2}, {3, 4}},
		},
		{
			name: "3",
			mat:  [][]int{{1, 2}, {3, 4}},
			r:    4,
			c:    1,
			res:  [][]int{{1}, {2}, {3}, {4}},
		},
		{
			name: "4",
			mat:  [][]int{{1, 2}, {3, 4}},
			r:    2,
			c:    4,
			res:  [][]int{{1, 2}, {3, 4}},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := matrixReshape(c.mat, c.r, c.c)
			assert.Equal(t, c.res, res)
		})
	}
}

func matrixReshape(mat [][]int, r int, c int) [][]int {
	nums := make([]int, 0, len(mat)*len(mat[0]))
	if cap(nums) != r*c {
		return mat
	}
	for i := range mat {
		for j := range mat[i] {
			nums = append(nums, mat[i][j])
		}
	}

	k := 0
	res := make([][]int, r)
	for i := range res {
		res[i] = make([]int, c)
		for j := range res[i] {
			res[i][j] = nums[k]
			k++
		}
	}

	return res
}
