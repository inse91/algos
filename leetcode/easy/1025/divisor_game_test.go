package _1025_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/divisor-game/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  bool
	}{
		{
			name: "1",
			n:    2,
			res:  true,
		},
		{
			name: "2",
			n:    3,
			res:  false,
		},
		{
			name: "3",
			n:    4,
			res:  true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, divisorGame(c.n))
		})
	}
}

func divisorGame(n int) bool {
	return n%2 == 0
}
