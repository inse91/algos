package _3516_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-closest-person

func Test(t *testing.T) {
	cases := []struct {
		name    string
		x, y, z int
		res     int
	}{
		{
			name: "1",
			x:    2,
			y:    7,
			z:    4,
			res:  1,
		},
		{
			name: "2",
			x:    2,
			y:    5,
			z:    6,
			res:  2,
		},
		{
			name: "3",
			x:    1,
			y:    5,
			z:    3,
			res:  0,
		},
		{
			name: "4",
			x:    14,
			y:    28,
			z:    4,
			res:  1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findClosest(c.x, c.y, c.z))
		})
	}
}

func findClosest(x int, y int, z int) int {
	r1 := z - x
	r2 := z - y
	if abs(r1) == abs(r2) {
		return 0
	}

	if abs(r1) > abs(r2) {
		return 2
	}

	return 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
