package _3360_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/stone-removal-game

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  bool
	}{
		{
			name: "1",
			n:    12,
			res:  true,
		},
		{
			name: "2",
			n:    1,
			res:  false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canAliceWin(c.n))
		})
	}
}

func canAliceWin(n int) bool {
	var (
		aliceTurn bool
		cur       = 10
	)

	for {
		if n < cur {
			return aliceTurn
		}

		n -= cur
		cur -= 1
		aliceTurn = !aliceTurn
	}
}
