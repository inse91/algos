package _3162_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-number-of-good-pairs-i

func Test(t *testing.T) {
	cases := []struct {
		name  string
		nums1 []int
		nums2 []int
		k     int
		res   int
	}{
		{
			name:  "1",
			nums1: []int{1, 3, 4},
			nums2: []int{1, 3, 4},
			k:     1,
			res:   5,
		},
		{
			name:  "2",
			nums1: []int{1, 2, 4, 12},
			nums2: []int{2, 4},
			k:     3,
			res:   2,
		},
		{
			name:  "13",
			nums1: []int{10, 25, 12},
			nums2: []int{1, 3, 5, 6},
			k:     2,
			res:   5,
		},
		{
			name:  "23",
			nums1: []int{10, 25, 12},
			nums2: []int{1, 3, 5, 6},
			k:     1,
			res:   7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfPairs(c.nums1, c.nums2, c.k))
		})
	}
}

func numberOfPairs(nums1 []int, nums2 []int, k int) int {
	m2 := make(map[int]int, len(nums2)/2)

	for _, num := range nums2 {
		m2[num*k]++
	}

	var res int
	for _, num := range nums1 {
		for i := 1; i <= num; i++ {
			if num%i != 0 {
				continue
			}

			cand := num / i
			res += m2[cand]
		}
	}

	return res
}
