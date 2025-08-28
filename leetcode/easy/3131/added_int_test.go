package _3131_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-integer-added-to-array-i

func Test(t *testing.T) {
	cases := []struct {
		name         string
		nums1, nums2 []int
		res          int
	}{
		{
			name:  "1",
			nums1: []int{2, 6, 4},
			nums2: []int{9, 5, 7},
			res:   3,
		},
		{
			name:  "2",
			nums1: []int{1, 1, 1},
			nums2: []int{1, 1, 1},
			res:   0,
		},
		{
			name:  "3",
			nums1: []int{10},
			nums2: []int{5},
			res:   -5,
		},
		{
			name:  "13",
			nums1: []int{2},
			nums2: []int{5},
			res:   3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, addedInteger(c.nums1, c.nums2))
		})
	}
}

func addedInteger(nums1 []int, nums2 []int) int {
	if len(nums1) == 1 {
		return nums2[0] - nums1[0]
	}

	m1 := slices.Min(nums1)
	m2 := slices.Min(nums2)

	return m2 - m1
}
