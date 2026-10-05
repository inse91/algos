package _3206_test

// https://leetcode.com/problems/alternating-groups-i

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		colors []int
		res    int
	}{
		{
			name:   "1",
			colors: []int{1, 1, 1},
			res:    0,
		},
		{
			name:   "2",
			colors: []int{0, 1, 0, 0, 1},
			res:    3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfAlternatingGroups(c.colors))
		})
	}
}

func numberOfAlternatingGroups(colors []int) int {
	if len(colors) < 3 {
		panic("not enough colors")
	}

	var (
		prv     = colors[len(colors)-1]
		nxt     = colors[1]
		nextIdx int
		c       int
	)

	for i, clr := range colors {
		switch clr {
		case 0:
			if prv == 1 && nxt == 1 {
				c++
			}
		case 1:
			if prv == 0 && nxt == 0 {
				c++
			}
		}

		nextIdx = i + 2
		if nextIdx >= len(colors) {
			nextIdx = 0
		}

		nxt = colors[nextIdx]
		prv = clr
	}

	return c
}
