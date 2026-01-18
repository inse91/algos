package _3046_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/split-the-array

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1, 1, 2, 2, 3, 4},
			res:  true,
		},
		{
			name: "2",
			nums: []int{1, 1, 1, 1},
			res:  false,
		},
		{
			name: "62",
			nums: []int{6, 1, 3, 1, 1, 8, 9, 2},
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isPossibleToSplit(c.nums))
		})
	}
}

func isPossibleToSplit(nums []int) bool {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
		if m[num] > 2 {
			return false
		}
	}

	return true
}
