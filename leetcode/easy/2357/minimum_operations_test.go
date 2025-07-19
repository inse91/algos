package _2357_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 5, 0, 3, 5},
			res:  3,
		},
		{
			name: "2",
			nums: []int{0},
			res:  0,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, minimumOperations(c.nums))
		})
	}
}

func minimumOperations(nums []int) int {
	m := make(map[int]struct{}, len(nums)/2)
	for _, v := range nums {
		if v == 0 {
			continue
		}

		m[v] = struct{}{}
	}

	return len(m)
}
