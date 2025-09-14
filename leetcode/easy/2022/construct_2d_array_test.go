package _2022_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/convert-1d-array-into-2d-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		m, n int
		res  [][]int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 4},
			m:    2,
			n:    2,
			res: [][]int{
				{1, 2},
				{3, 4},
			},
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4},
			m:    1,
			n:    4,
			res: [][]int{
				{1, 2, 3, 4},
			},
		},
		{
			name: "3",
			nums: []int{1, 2},
			m:    1,
			n:    1,
			res:  nil,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, construct2DArray(c.nums, c.m, c.n))
		})
	}
}

func construct2DArray(original []int, m int, n int) [][]int {
	if len(original) != m*n {
		return nil
	}

	res := make([][]int, m)
	var c int
	for i := range res {
		res[i] = make([]int, 0, n)
		for range n {
			res[i] = append(res[i], original[c])
			c++
		}
	}

	return res
}
