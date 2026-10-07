package _3222_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/find-the-winning-player-in-coin-game

func Test(t *testing.T) {
	cases := []struct {
		name string
		x, y int
		res  string
	}{
		{
			name: "1",
			x:    2,
			y:    7,
			res:  "Alice",
		},
		{
			name: "2",
			x:    4,
			y:    11,
			res:  "Bob",
		},
		{
			name: "3",
			x:    3,
			y:    17,
			res:  "Alice",
		},
		{
			name: "4",
			x:    6,
			y:    27,
			res:  "Bob",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, winningPlayer(
				c.x,
				c.y,
			))
		})
	}
}

func winningPlayer(x int, y int) string {
	var turn bool
	for {
		if x == 0 || y < 4 {
			break
		}

		x -= 1
		y -= 4
		turn = !turn
	}

	if turn {
		return "Alice"
	}

	return "Bob"
}
