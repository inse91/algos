package _2956_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-common-elements-between-two-arrays

func Test(t *testing.T) {
	cases := []struct {
		name         string
		nums1, nums2 []int
		res          []int
	}{
		{
			name:  "1",
			nums1: []int{2, 3, 2},
			nums2: []int{1, 2},
			res:   []int{2, 1},
		},
		{
			name:  "2",
			nums1: []int{4, 3, 2, 3, 1},
			nums2: []int{2, 2, 5, 2, 3, 6},
			res:   []int{3, 4},
		},
		{
			name:  "3",
			nums1: []int{3, 4, 2, 3},
			nums2: []int{1, 5},
			res:   []int{0, 0},
		},
		{
			name:  "509",
			nums1: []int{25, 8, 14, 61, 43, 24, 69, 8, 41, 81, 78, 27, 8, 7, 43, 8, 13, 62, 41, 24, 1, 31, 41, 69, 41, 78, 61},
			nums2: []int{44, 11, 9, 4, 100, 6, 8, 13, 81, 11, 48, 61, 43, 24, 24, 71, 54, 63, 53, 40, 41, 16, 41, 92, 89, 62, 66, 46, 41, 16},
			res:   []int{17, 11},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findIntersectionValues(c.nums1, c.nums2))
		})
	}
}

func findIntersectionValues(nums1 []int, nums2 []int) []int {
	var nums11 [101]bool
	for _, v := range nums1 {
		nums11[v] = true
	}

	var c1 int
	var nums22 [101]bool
	for _, v := range nums2 {
		nums22[v] = true
		if nums11[v] {
			c1++
		}
	}

	if c1 == 0 {
		return []int{0, 0}
	}

	var c2 int
	for _, v := range nums1 {
		if nums22[v] {
			c2++
		}
	}

	return []int{c2, c1}
}
