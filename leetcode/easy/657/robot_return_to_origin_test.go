package _657_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRobotReturnToOrigin(t *testing.T) {
	cases := []struct {
		name string
		m    string
		res  bool
	}{
		{
			name: "1",
			m:    "UD",
			res:  true,
		},
		{
			name: "2",
			m:    "LL",
			res:  false,
		},
		{
			name: "3",
			m:    "RL",
			res:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := judgeCircle(c.m)
			assert.Equal(t, c.res, res)
		})
	}
}

func judgeCircle(moves string) bool {
	if len(moves) == 0 {
		return true
	}
	var x, y int
	for _, r := range moves {
		switch r {
		case 'R':
			x++
		case 'L':
			x--
		case 'D':
			y++
		case 'U':
			y--
		}
	}

	return x == 0 && y == 0
}
