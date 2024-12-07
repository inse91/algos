package _1496_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/path-crossing/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		path string
		res  bool
	}{
		{
			name: "1",
			path: "NES",
			res:  false,
		},
		{
			name: "2",
			path: "NESWW",
			res:  true,
		},
		{
			name: "3",
			path: "SS",
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, isPathCrossing(c.path))
		})
	}
}

type boy struct {
	current [2]int
	history map[[2]int]struct{}
}

func (b *boy) takeStep(r rune) bool {
	point := [2]int{b.current[0], b.current[1]}
	switch r {
	case 'N':
		point[1]++
	case 'S':
		point[1]--
	case 'W':
		point[0]++
	case 'E':
		point[0]--
	}

	_, ok := b.history[point]
	if ok {
		return true
	}

	b.history[point] = struct{}{}
	b.current = point

	return false
}

func isPathCrossing(path string) bool {
	newBoy := boy{
		current: [2]int{0, 0},
		history: make(map[[2]int]struct{}, len(path)),
	}
	newBoy.history[newBoy.current] = struct{}{}

	for _, r := range path {
		if newBoy.takeStep(r) {
			return true
		}
	}

	return false
}
