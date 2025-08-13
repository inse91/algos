package _2643_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/row-with-maximum-ones

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums [][]int
		res  []int
	}{
		{
			name: "1",
			nums: [][]int{{0, 1}, {1, 0}},
			res:  []int{0, 1},
		},
		{
			name: "2",
			nums: [][]int{{0, 0, 0}, {0, 1, 1}},
			res:  []int{1, 2},
		},
		{
			name: "3",
			nums: [][]int{{0, 0}, {1, 1}, {0, 0}},
			res:  []int{1, 2},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, rowAndMaximumOnes(c.nums))
		})
	}
}

func rowAndMaximumOnes(mat [][]int) []int {
	var mx, mxIdx int
	for i, nums := range mat {
		var sum int
		for _, num := range nums {
			sum += num
		}

		if sum <= mx {
			continue
		}

		mx = sum
		mxIdx = i
	}

	return []int{mxIdx, mx}
}
