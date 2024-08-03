package _1207_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/unique-number-of-occurrences/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{1, 2, 2, 1, 1, 3},
			res:  true,
		},
		{
			name: "2",
			nums: []int{1, 2},
			res:  false,
		},
		{
			name: "3",
			nums: []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0},
			res:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, uniqueOccurrences(c.nums))
		})
	}
}

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for _, num := range arr {
		m[num]++
	}

	set := make(map[int]struct{})
	for _, v := range m {
		if _, ok := set[v]; ok {
			return false
		}
		set[v] = struct{}{}
	}

	return true
}
