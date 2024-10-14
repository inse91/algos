package _1460_test

import (
	"maps"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/make-two-arrays-equal-by-reversing-subarrays/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		target []int
		arr    []int
		res    bool
	}{
		{
			name:   "1",
			target: []int{1, 2, 3, 4},
			arr:    []int{2, 4, 1, 3},
			res:    true,
		},
		{
			name:   "2",
			target: []int{7},
			arr:    []int{7},
			res:    true,
		},
		{
			name:   "3",
			target: []int{3, 7, 9},
			arr:    []int{3, 7, 11},
			res:    false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canBeEqual(c.target, c.arr))
		})
	}
}

func canBeEqual(target []int, arr []int) bool {
	m1 := make(map[int]int, len(target))
	m2 := make(map[int]int, len(arr))

	for i := range target {
		m1[target[i]]++
		m2[arr[i]]++
	}

	return maps.Equal(m1, m2)
}
