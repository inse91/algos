package _724_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPivotIdx(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 7, 3, 6, 5, 6},
			res:  3,
		},
		{
			name: "2",
			nums: []int{1, 2, 3},
			res:  -1,
		},
		{
			name: "3",
			nums: []int{2, 1, -1},
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			res := pivotIndex(c.nums)
			assert.Equal(t, c.res, res)
		})
	}
}

func pivotIndex(nums []int) int {
	var r int
	for i := range nums {
		r += nums[i]
	}

	var l int
	var prev int
	for i, num := range nums {
		r -= num
		l += prev
		if r == l {
			return i
		}
		prev = num
	}

	return -1
}
