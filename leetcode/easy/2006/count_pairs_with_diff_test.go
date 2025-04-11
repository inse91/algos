package _2006_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/count-number-of-pairs-with-absolute-difference-k/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 1},
			k:    1,
			res:  4,
		},
		{
			name: "2",
			nums: []int{1, 3},
			k:    3,
			res:  0,
		},
		{
			name: "3",
			nums: []int{3, 2, 1, 5, 4},
			k:    2,
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countKDifference(c.nums, c.k))
		})
	}
}

func countKDifference(nums []int, k int) int {
	var res int
	m := map[int]int{}
	for _, n := range nums {
		m[n]++
	}

	for num, q := range m {
		pair := num - k
		if v, ok := m[pair]; ok {
			res += v * q
		}
	}

	return res
}
