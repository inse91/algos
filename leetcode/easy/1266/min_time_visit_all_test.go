package _1266_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		pts  [][]int
		res  int
	}{
		{name: "1", pts: [][]int{{1, 1}, {3, 4}, {-1, 0}}, res: 7},
		{name: "2", pts: [][]int{{3, 2}, {-2, 2}}, res: 5},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minTimeToVisitAllPoints(c.pts))
		})
	}
}

func minTimeToVisitAllPoints(points [][]int) int {
	var res int
	for i := 1; i < len(points); i++ {
		p := points[i]
		pp := points[i-1]
		res += dist(
			[2]int{p[0], p[1]},
			[2]int{pp[0], pp[1]},
		)
	}

	return res
}
func dist(a, b [2]int) int {
	x := abs(a[0], b[0])
	y := abs(a[1], b[1])

	return max(x, y)
}

func abs(a, b int) int {
	res := a - b
	if res < 0 {
		return -res
	}

	return res
}
