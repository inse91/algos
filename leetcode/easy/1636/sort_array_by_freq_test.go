package _1636_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-array-by-increasing-frequency/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  []int
	}{
		{
			name: "1",
			nums: []int{1, 1, 2, 2, 2, 3},
			res:  []int{3, 1, 1, 2, 2, 2},
		},
		{
			name: "2",
			nums: []int{2, 3, 1, 3, 2},
			res:  []int{1, 3, 3, 2, 2},
		},
		{
			name: "3",
			nums: []int{-1, 1, -6, 4, 5, -6, 1, 4, 1},
			res:  []int{5, -1, 4, 4, -6, -6, 1, 1, 1},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, frequencySort(c.nums))
		})
	}
}

func frequencySort(nums []int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	m := make(map[int][]int, len(freqMap))
	for k, v := range freqMap {
		m[v] = append(m[v], k)
	}

	freqs := make([]int, 0, len(m))
	for k := range m {
		freqs = append(freqs, k)
	}
	slices.Sort(freqs)

	res := make([]int, 0, len(nums))
	for _, f := range freqs {
		q := m[f]
		slices.SortFunc(q, func(a, b int) int {
			if a < b {
				return 1
			}

			return -1
		})

		for _, _q := range q {
			for i := 0; i < f; i++ {
				res = append(res, _q)
			}
		}
	}

	return res
}
