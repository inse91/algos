package _1122_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		a1   []int
		a2   []int
		res  []int
	}{
		{
			name: "1",
			a1:   []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19},
			a2:   []int{2, 1, 4, 3, 9, 6},
			res:  []int{2, 2, 2, 1, 4, 3, 3, 9, 6, 7, 19},
		},
		{
			name: "2",
			a1:   []int{28, 6, 22, 8, 44, 17},
			a2:   []int{22, 28, 8, 6},
			res:  []int{22, 28, 8, 6, 17, 44},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.True(t, slices.Equal(c.res, relativeSortArray(c.a1, c.a2)))
		})
	}
}

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	r := make(map[int]struct{})

	result := make([]int, 0, len(arr1))
	rest := make([]int, 0)

	for _, n2 := range arr2 {
		r[n2] = struct{}{}
	}

	for _, n1 := range arr1 {
		_, ok := r[n1]
		if !ok {
			rest = append(rest, n1)
			continue
		}
		m[n1]++
	}

	slices.Sort(rest)

	for _, n2 := range arr2 {
		times := m[n2]
		for i := 0; i < times; i++ {
			result = append(result, n2)
		}
	}

	result = append(result, rest...)

	return result
}
