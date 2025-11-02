package _2778_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-squares-of-special-elements

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 4},
			res:  21,
		},
		{
			name: "2",
			nums: []int{2, 7, 1, 19, 18, 3},
			res:  63,
		},
		{
			name: "3",
			nums: []int{},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOfSquares(c.nums))
		})
	}
}

func sumOfSquares(nums []int) int {
	var (
		res int
		ln  = len(nums)
	)
	for i, num := range nums {
		i++
		if ln%i == 0 {
			res += num * num
		}
	}

	return res
}
