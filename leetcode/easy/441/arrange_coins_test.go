package _441_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrangeCoins(t *testing.T) {
	cases := []struct {
		name string
		in   int
		res  int
	}{
		{
			name: "1",
			in:   5,
			res:  2,
		},
		{
			name: "2",
			in:   8,
			res:  3,
		},
		{
			name: "3",
			in:   1,
			res:  1,
		},
		{
			name: "4",
			in:   0,
			res:  0,
		},
		{
			name: "5",
			in:   3,
			res:  2,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := arrangeCoins(c.in)
			assert.Equal(t, c.res, res)
		})
	}
}

func arrangeCoins(n int) (stairs int) {
	coinsToBuildOneLine := 1
	for {
		n -= coinsToBuildOneLine
		if n < 0 {
			break
		}
		coinsToBuildOneLine++
		stairs++
	}
	return stairs
}
