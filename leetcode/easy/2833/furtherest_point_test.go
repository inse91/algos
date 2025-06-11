package _833

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/furthest-point-from-origin

func Test(t *testing.T) {
	cases := []struct {
		name  string
		moves string
		res   int
	}{
		{
			name:  "1",
			moves: "L_RL__R",
			res:   3,
		},
		{
			name:  "2",
			moves: "_R__LL_",
			res:   5,
		},
		{
			name:  "3",
			moves: "_______",
			res:   7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, furthestDistanceFromOrigin(c.moves))
		})
	}
}

func furthestDistanceFromOrigin(moves string) int {
	var ls, rs int
	for _, mv := range moves {
		switch mv {
		case 'L':
			ls += 1
		case 'R':
			rs += 1
		default:
		}
	}

	blancs := len(moves) - rs - ls
	if ls > rs {
		return ls - rs + blancs
	}

	return rs - ls + blancs
}
