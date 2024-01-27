package _455_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func TestAssignCookies(t *testing.T) {
	cases := []struct {
		name string
		g    []int
		s    []int
		res  int
	}{
		{
			name: "1",
			g:    []int{1, 2, 3},
			s:    []int{1, 1},
			res:  1,
		},
		{
			name: "2",
			g:    []int{1, 2},
			s:    []int{1, 2, 3},
			res:  2,
		},
		{
			name: "3",
			g:    []int{10, 9, 8, 7},
			s:    []int{5, 6, 7, 8},
			res:  2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findContentChildren(c.g, c.s)
			assert.Equal(t, c.res, res)
		})
	}
}

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	slices.Sort(g)
	slices.Sort(s)

	var res int
	for i, j := 0, 0; i < len(g) && j < len(s); j++ {
		if g[i] > s[j] {
			continue
		}
		i++
		res++
	}
	return res
}
