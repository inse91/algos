package _349__test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntersection(t *testing.T) {
	cases := []struct {
		name string
		in1  []int
		in2  []int
		res  []int
	}{
		{
			name: "1",
			in1:  []int{1, 2, 2, 1},
			in2:  []int{2, 2},
			res:  []int{2},
		},
		{
			name: "2",
			in1:  []int{4, 9, 5},
			in2:  []int{9, 4, 9, 8, 4},
			res:  []int{9, 4},
		},
		{
			name: "3",
			in1:  []int{},
			in2:  []int{1},
			res:  []int{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := intersection(c.in1, c.in2)
			assert.ElementsMatch(t, c.res, res)
		})
	}
}

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, num := range nums1 {
		m[num] = false
	}

	for _, num := range nums2 {
		_, ok := m[num]
		if !ok {
			continue
		}
		m[num] = true
	}

	res := make([]int, 0, len(m)/2)
	for k, v := range m {
		if v {
			res = append(res, k)
		}
	}

	return res
}
