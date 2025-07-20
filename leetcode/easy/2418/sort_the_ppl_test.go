package _2418_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-the-people

func Test(t *testing.T) {
	cases := []struct {
		name  string
		names []string
		h     []int
		res   []string
	}{
		{
			name:  "1",
			names: []string{"Mary", "John", "Emma"},
			h:     []int{180, 165, 170},
			res:   []string{"Mary", "Emma", "John"},
		},
		{
			name:  "2",
			names: []string{"Alice", "Bob", "Bob"},
			h:     []int{155, 185, 150},
			res:   []string{"Bob", "Alice", "Bob"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, sortPeople(c.names, c.h))
		})
	}
}

func sortPeople(names []string, heights []int) []string {
	m := make(map[int]string, len(names))
	for i, h := range heights {
		m[h] = names[i]
	}

	slices.Sort(heights)
	slices.Reverse(heights)
	for i, height := range heights {
		names[i] = m[height]
	}

	return names
}
