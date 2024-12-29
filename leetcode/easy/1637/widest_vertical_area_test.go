package _1637_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		pts  [][]int
		res  int
	}{
		{
			name: "1",
			pts:  [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}},
			res:  1,
		},
		{
			name: "2",
			pts:  [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}},
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxWidthOfVerticalArea(c.pts))
		})
	}
}

func maxWidthOfVerticalArea(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		if a[0] < b[0] {
			return -1
		}

		return 1
	})

	ln := -1
	for i := 1; i < len(points); i++ {
		ln = max(ln, points[i][0]-points[i-1][0])
	}

	return ln
}
