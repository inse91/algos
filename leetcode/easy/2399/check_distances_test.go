package _2399_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-distances-between-same-letters

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		nums []int
		res  bool
	}{
		{
			name: "1",
			s:    "abaccb",
			nums: []int{1, 3, 0, 5, 0, 0, 0, 0, 0, 0},
			res:  true,
		},
		{
			name: "2",
			s:    "aa",
			nums: []int{1, 0, 0, 0, 0, 0, 0, 0, 0},
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, checkDistances(c.s, c.nums))
		})
	}
}

func checkDistances(s string, distances []int) bool {
	m := make(map[rune]int, 26)
	for i, r := range s {
		key := r - 'a'
		if _, ok := m[key]; !ok {
			m[key] = i
			continue
		}

		m[key] = i - m[key] - 1
	}

	for i, d := range distances {
		v, ok := m[rune(i)]
		if !ok {
			continue
		}

		if v != d {
			return false
		}
	}

	return true
}
