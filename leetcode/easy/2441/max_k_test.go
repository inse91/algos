package _2441_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/largest-positive-integer-that-exists-with-its-negative

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{-1, 2, -3, 3},
			res:  3,
		},
		{
			name: "2",
			nums: []int{-1, 10, 6, 7, -7, 1},
			res:  7,
		},
		{
			name: "3",
			nums: []int{-10, 8, 6, 7, -2, -3},
			res:  -1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findMaxK(c.nums))
		})
	}
}

func findMaxK(nums []int) int {
	res := -1
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}

	for k := range m {
		if k <= 0 {
			continue
		}

		_, ok := m[-k]
		if !ok {
			continue
		}

		res = max(res, k)
	}

	return res
}
