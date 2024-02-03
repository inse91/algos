package _496_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestNextGreaterElem(t *testing.T) {
	cases := []struct {
		name  string
		nums1 []int
		nums2 []int
		res   []int
	}{
		{
			name:  "1",
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			res:   []int{-1, 3, -1},
		},
		{
			name:  "2",
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			res:   []int{3, -1},
		},
		{
			name:  "3",
			nums1: []int{1, 3, 5, 2, 4},
			nums2: []int{6, 5, 4, 3, 2, 1, 7},
			res:   []int{7, 7, 7, 7, 7},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := nextGreaterElement(c.nums1, c.nums2)
			assert.Equal(t, c.res, res)
		})
	}
}

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := make([]int, 0, len(nums1))
outer:
	for _, num := range nums1 {
		nums2idx := slices.Index(nums2, num)
		if nums2idx == -1 {
			res = append(res, -1)
			continue
		}
		for _, nums2num := range nums2[nums2idx:] {
			if nums2num > num {
				res = append(res, nums2num)
				continue outer
			}
		}
		res = append(res, -1)
	}

	return res
}
