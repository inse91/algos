package _1779_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/

func Test(t *testing.T) {
	cases := []struct {
		name string
		x    int
		y    int
		pts  [][]int
		res  int
	}{
		{
			name: "1",
			x:    3,
			y:    4,
			pts:  [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}},
			res:  2,
		},
		{
			name: "2",
			x:    3,
			y:    4,
			pts:  [][]int{{3, 4}},
			res:  0,
		},
		{
			name: "3",
			x:    3,
			y:    4,
			pts:  [][]int{{2, 3}},
			res:  -1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, nearestValidPoint(c.x, c.y, c.pts))
		})
	}
}

func nearestValidPoint(x int, y int, points [][]int) int {
	idx := -1
	var mn int = 10e4 + 1
	for i, point := range points {
		xP, yP := point[0], point[1]
		if xP == x {
			gap := abs(yP - y)
			if gap >= mn {
				continue
			}

			mn = gap
			idx = i
		}

		if yP == y {
			gap := abs(xP - x)
			if gap >= mn {
				continue
			}

			mn = gap
			idx = i
		}
	}

	return idx
}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}
