package _1725_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/number-of-rectangles-that-can-form-the-largest-square/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		rects [][]int
		res   int
	}{
		{
			name:  "1",
			rects: [][]int{{5, 8}, {3, 9}, {5, 12}, {16, 5}},
			res:   3,
		},
		{
			name:  "2",
			rects: [][]int{{2, 3}, {3, 7}, {4, 3}, {3, 7}},
			res:   3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, countGoodRectangles(c.rects))
		})
	}
}

func countGoodRectangles(rectangles [][]int) int {
	var res int
	cur := 1

	for _, r := range rectangles {
		mx := min(r[0], r[1])
		if mx == cur {
			res++
			continue
		}

		if mx > cur {
			res = 1
			cur = mx
		}
	}

	return res
}
