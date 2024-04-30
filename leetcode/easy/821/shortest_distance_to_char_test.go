package _821_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShortestDistanceToChar(t *testing.T) {
	cases := []struct {
		name string
		s    string
		c    byte
		res  []int
	}{
		{
			name: "1",
			s:    "loveleetcode",
			c:    'e',
			res:  []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0},
		},
		{
			name: "2",
			s:    "aaab",
			c:    'b',
			res:  []int{3, 2, 1, 0},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, shortestToChar(c.s, c.c))
		})
	}
}

func shortestToChar(s string, c byte) []int {
	cIdxes := make(map[int]struct{})
	for i, r := range s {
		if byte(r) == c {
			cIdxes[i] = struct{}{}
		}
	}
	ci := make([]int, 0, len(cIdxes))
	for k := range cIdxes {
		ci = append(ci, k)
	}

	var sl []int
	res := make([]int, 0, len(s))
	for i := range s {
		sl = make([]int, 0, len(cIdxes))
		for _, j := range ci {
			dist := abs(i, j)
			sl = append(sl, dist)
		}
		if len(sl) == 0 {
			panic("no can do")
		}
		if len(sl) == 1 {
			res = append(res, sl[0])
		} else {
			mn := findMin(sl)
			res = append(res, mn)
		}

		clear(sl)
	}

	return res
}

func findMin(sl []int) int {
	if len(sl) == 0 {
		return -1
	}
	mn := sl[0]
	for _, n := range sl {
		if n == 0 {
			return 0
		}
		mn = min(mn, n)
	}

	return mn
}

func abs(a, b int) int {
	r := a - b
	if r < 0 {
		return -r
	}

	return r
}
