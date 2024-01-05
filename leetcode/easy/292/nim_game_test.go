package _292_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNimGame(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  bool
	}{
		{
			name: "1",
			in:   4,
			res:  false,
		},
		{
			name: "2",
			in:   1,
			res:  true,
		},
		{
			name: "3",
			in:   2,
			res:  true,
		},
		{
			name: "4",
			in:   3,
			res:  true,
		},
		{
			name: "5",
			in:   5,
			res:  true,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := canWinNim(c.in)
			assert.Equal(t, c.res, res)
		})

	}
}

func canWinNim(n int) bool {
	return n%4 != 0
}
