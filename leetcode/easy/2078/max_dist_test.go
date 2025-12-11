package _2078_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/two-furthest-houses-with-different-colors

func Test(t *testing.T) {
	cases := []struct {
		name   string
		colors []int
		res    int
	}{
		{
			name:   "1",
			colors: []int{1, 1, 1, 6, 1, 1, 1},
			res:    3,
		},
		{
			name:   "2",
			colors: []int{1, 8, 3, 8, 3},
			res:    4,
		},
		{
			name:   "3",
			colors: []int{0, 1},
			res:    1,
		},
		{
			name:   "82",
			colors: []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 19, 19, 6, 6},
			res:    10,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxDistance(c.colors))
		})
	}
}

func maxDistance(colors []int) int {
	var (
		first      = colors[0]
		ln         = len(colors)
		last       = colors[ln-1]
		res        int
		foundRight bool
		foundLeft  bool
	)

	for i := range ln {
		if foundLeft && foundRight {
			break
		}

		r := colors[ln-i-1]
		if first != r && !foundRight {
			res = max(res, ln-i)
			foundRight = true
		}

		l := colors[i]
		if last != l && !foundLeft {
			res = max(res, ln-i)
			foundLeft = true
		}
	}

	return res - 1
}
