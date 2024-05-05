package _766_test

// https://leetcode.com/problems/toeplitz-matrix/submissions/1249216684/

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtx  [][]int
		res  bool
	}{
		{
			name: "1",
			mtx:  [][]int{{1, 2, 3, 4}, {5, 1, 2, 3}, {9, 5, 1, 2}},
			res:  true,
		},
		{
			name: "2",
			mtx:  [][]int{{1, 2}, {2, 2}},
			res:  false,
		},
		{
			name: "364",
			mtx:  [][]int{{11, 74, 0, 93}, {40, 11, 74, 7}},
			res:  false,
		},
		{
			name: "364_true",
			mtx:  [][]int{{11, 74, 7, 93}, {40, 11, 74, 7}},
			res:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isToeplitzMatrix(c.mtx))
		})
	}
}

func isToeplitzMatrix(matrix [][]int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	if len(matrix) == 1 || len(matrix[0]) == 1 {
		return true
	}

	// center
	d := matrix[0][0]
	for i := 1; i < len(matrix) && i < len(matrix[0]); i++ {
		v := matrix[i][i]
		if v != d {
			return false
		}
	}

	for q := 1; q < len(matrix[0])-1; q++ {
		d = matrix[0][q]
		for i := 1; i < len(matrix) && q+i < len(matrix[0]); i++ {
			v := matrix[i][q+i]
			if v != d {
				return false
			}
		}
	}

	for q := 1; q < len(matrix)-1; q++ {
		d = matrix[q][0]
		for i := 1; i < len(matrix[0]) && q+i < len(matrix); i++ {
			v := matrix[q+i][i]
			if v != d {
				return false
			}
		}
	}

	return true
}
