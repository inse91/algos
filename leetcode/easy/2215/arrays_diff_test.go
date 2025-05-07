package _2215_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-difference-of-two-arrays/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		nums1 []int
		nums2 []int
		res   [][]int
	}{
		{
			name:  "1",
			nums1: []int{1, 2, 3},
			nums2: []int{2, 4, 6},
			res:   [][]int{{1, 3}, {4, 6}},
		},
		{
			name:  "2",
			nums1: []int{1, 2, 3, 3},
			nums2: []int{1, 1, 2, 2},
			res:   [][]int{{3}, {}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			diff := findDifference(c.nums1, c.nums2)
			assert.ElementsMatch(t, c.res[0], diff[0])
			assert.ElementsMatch(t, c.res[1], diff[1])
		})
	}
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	res := [][]int{{}, {}}
	m1 := map[int]struct{}{}
	for _, n := range nums1 {
		m1[n] = struct{}{}
	}

	m2 := map[int]struct{}{}
	for _, n := range nums2 {
		m2[n] = struct{}{}
	}

	for n := range m1 {
		if _, ok := m2[n]; !ok {
			res[0] = append(res[0], n)
		}
	}

	for n := range m2 {
		if _, ok := m1[n]; !ok {
			res[1] = append(res[1], n)
		}
	}

	return res
}
