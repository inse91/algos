package _2154_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/keep-multiplying-found-values-by-two

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		orig int
		res  int
	}{
		{
			name: "1",
			nums: []int{5, 3, 6, 1, 12},
			orig: 3,
			res:  24,
		},
		{
			name: "2",
			nums: []int{2, 7, 9},
			orig: 4,
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findFinalValue(c.nums, c.orig))
		})
	}
}

func findFinalValue(nums []int, original int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	for {
		if _, ok := m[original]; !ok {
			return original
		}

		original *= 2
	}
}
