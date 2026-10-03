package _3238_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-number-of-winning-players

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		pick [][]int
		res  int
	}{
		{
			name: "1",
			n:    4,
			pick: [][]int{
				{0, 0},
				{1, 0},
				{1, 0},
				{2, 1},
				{2, 1},
				{2, 0},
			},
			res: 2,
		},
		{
			name: "2",
			n:    5,
			pick: [][]int{
				{1, 1},
				{1, 2},
				{1, 3},
				{1, 4},
			},
			res: 0,
		},
		{
			name: "3",
			n:    5,
			pick: [][]int{
				{1, 1},
				{2, 4},
				{2, 4},
				{2, 4},
			},
			res: 1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, winningPlayerCount(c.n, c.pick))
		})
	}
}

func winningPlayerCount(n int, picks [][]int) int {
	m := make(map[int]map[int]int, n)
	for i := range n {
		m[i] = make(map[int]int)
	}

	for _, pick := range picks {
		m[pick[0]][pick[1]] += 1
	}

	var c int
	for player, balls := range m {
		for _, num := range balls {
			if num > player {
				c++

				break
			}
		}
	}

	return c
}
