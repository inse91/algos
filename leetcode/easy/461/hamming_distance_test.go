package _461__test

import (
	"github.com/stretchr/testify/assert"
	"math/bits"
	"testing"
)

func TestHammingDistance(t *testing.T) {
	cases := []struct {
		name string
		x    int
		y    int
		res  int
	}{
		{
			name: "1",
			x:    1,
			y:    4,
			res:  2,
		},
		{
			name: "2",
			x:    3,
			y:    1,
			res:  1,
		},
		{
			name: "3",
			x:    1,
			y:    1,
			res:  0,
		},
		{
			name: "4",
			x:    0,
			y:    1,
			res:  1,
		},
		{
			name: "5",
			x:    0,
			y:    0,
			res:  0,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := hammingDistance(c.x, c.y)
			assert.Equal(t, c.res, res)
		})
	}
}
func hammingDistance(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
