package _1995_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-special-quadruplets

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 3, 6},
			res:  1,
		},
		{
			name: "2",
			nums: []int{3, 3, 6, 4, 5},
			res:  0,
		},
		{
			name: "3",
			nums: []int{1, 1, 1, 3, 5},
			res:  4,
		},
		{
			name: "39",
			nums: []int{28, 8, 49, 85, 37, 90, 20, 8},
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countQuadruplets(c.nums))
		})
	}
}

func countQuadruplets(nums []int) int {
	var res int

	ln := len(nums)
	for i, a := range nums {
		for j := i + 1; j < ln; j++ {
			b := nums[j]
			for k := j + 1; k < ln; k++ {
				c := nums[k]
				for p := k + 1; p < ln; p++ {
					d := nums[p]
					sm := a + b + c
					if sm == d {
						res++
					}
				}
			}
		}
	}

	return res
}
