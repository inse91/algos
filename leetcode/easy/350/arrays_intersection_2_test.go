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
			res:  []int{2, 2},
		},
		{
			name: "2",
			in1:  []int{4, 9, 5},
			in2:  []int{9, 4, 9, 8, 4},
			res:  []int{9, 4},
		},
		{
			name: "3",
			in1:  []int{4, 4, 9, 5},
			in2:  []int{9, 4, 9, 8, 4},
			res:  []int{9, 4, 4},
		},
		{
			name: "4",
			in1:  []int{},
			in2:  []int{1},
			res:  []int{},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := intersect(c.in1, c.in2)
			assert.ElementsMatch(t, c.res, res)
		})
	}
}

func count(in []int) map[int]int {
	m := make(map[int]int)
	for _, num := range in {
		if _, ok := m[num]; !ok {
			m[num] = 1
			continue
		}
		m[num]++
	}
	return m
}

func intersect(nums1 []int, nums2 []int) []int {
	m1 := count(nums1)
	m2 := count(nums2)

	res := make([]int, 0, len(m1)/2)
	for k1, v1 := range m1 {
		v2, ok := m2[k1]
		if !ok {
			continue
		}
		for i := 0; i < min(v1, v2); i++ {
			res = append(res, k1)
		}
	}

	return res
}
