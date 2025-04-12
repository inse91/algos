package _2032_test

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-out-of-three/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		nums1 []int
		nums2 []int
		nums3 []int
		res   []int
	}{
		{
			name:  "1",
			nums1: []int{1, 1, 3, 2},
			nums2: []int{2, 3},
			nums3: []int{3},
			res:   []int{3, 2},
		},
		{
			name:  "2",
			nums1: []int{3, 1},
			nums2: []int{2, 3},
			nums3: []int{1, 2},
			res:   []int{2, 3, 1},
		},
		{
			name:  "3",
			nums1: []int{1, 2, 2},
			nums2: []int{4, 3, 3},
			nums3: []int{5},
			res:   []int{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, twoOutOfThree(c.nums1, c.nums2, c.nums3))
		})
	}
}

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	out := make(map[int]struct{})

	q1 := make(map[int]struct{}, len(nums1))
	for _, n := range nums1 {
		q1[n] = struct{}{}
	}

	q2 := make(map[int]struct{}, len(nums2))
	for _, n := range nums2 {
		if _, ok := q1[n]; ok {
			out[n] = struct{}{}
		}
		q2[n] = struct{}{}
	}

	for _, n := range nums3 {
		if _, ok := q1[n]; ok {
			out[n] = struct{}{}
			continue
		}
		if _, ok := q2[n]; ok {
			out[n] = struct{}{}
		}
	}

	res := make([]int, 0, len(out))
	for v := range maps.Keys(out) {
		res = append(res, v)
	}

	return res
}
