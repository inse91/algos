package _3375_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{5, 2, 5, 4, 5},
			k:    2,
			res:  2,
		},
		{
			name: "2",
			nums: []int{2, 1, 2},
			k:    2,
			res:  -1,
		},
		{
			name: "3",
			nums: []int{9, 7, 5, 3},
			k:    1,
			res:  4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minOperations(c.nums, c.k))
		})
	}
}

func minOperations(nums []int, k int) int {
	m := make(map[int]struct{}, len(nums)/2)

	var res int
	for _, num := range nums {
		if num < k {
			return -1
		}

		if num == k {
			continue
		}

		if _, ok := m[num]; ok {
			continue
		}

		m[num] = struct{}{}
		res++
	}

	return res
}
