package _2570_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values

func Test(t *testing.T) {
	cases := []struct {
		name  string
		nums1 [][]int
		nums2 [][]int
		res   [][]int
	}{
		{
			name: "1",
			nums1: [][]int{
				{1, 2},
				{2, 3},
				{4, 5},
			},
			nums2: [][]int{
				{1, 4},
				{3, 2},
				{4, 1},
			},
			res: [][]int{
				{1, 6},
				{2, 3},
				{3, 2},
				{4, 6},
			},
		},
		{
			name: "2",
			nums1: [][]int{
				{2, 4},
				{3, 6},
				{5, 5},
			},
			nums2: [][]int{
				{1, 3},
				{4, 3},
			},
			res: [][]int{
				{1, 3},
				{2, 4},
				{3, 6},
				{4, 3},
				{5, 5},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, mergeArrays(c.nums1, c.nums2))
		})
	}
}

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int]int, len(nums1)+len(nums2))
	for _, nums := range append(nums1, nums2...) {
		m[nums[0]] += nums[1]
	}

	res := make([][]int, 0, len(m))
	for k, v := range m {
		res = append(res, []int{k, v})
	}

	slices.SortFunc(res, func(a, b []int) int {
		return a[0] - b[0]
	})

	return res
}
