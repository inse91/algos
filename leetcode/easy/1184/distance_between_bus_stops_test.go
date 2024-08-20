package _1184_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/distance-between-bus-stops/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		dist  []int
		start int
		dest  int
		res   int
	}{
		{
			name:  "1",
			dist:  []int{1, 2, 3, 4},
			start: 0,
			dest:  1,
			res:   1,
		},
		{
			name:  "1_1",
			dist:  []int{1, 2, 3, 4},
			start: 2,
			dest:  0,
			res:   3,
		},
		{
			name:  "1_2",
			dist:  []int{1, 2, 3, 4},
			start: 3,
			dest:  0,
			res:   4,
		},
		{
			name:  "1_3",
			dist:  []int{1, 2, 3, 4},
			start: 0,
			dest:  3,
			res:   4,
		},
		{
			name:  "2",
			dist:  []int{1, 2, 3, 4},
			start: 0,
			dest:  2,
			res:   3,
		},
		{
			name:  "3",
			dist:  []int{1, 2, 3, 4},
			start: 0,
			dest:  3,
			res:   4,
		},
		{
			name:  "23",
			dist:  []int{8, 11, 6, 7, 10, 11, 2},
			start: 0,
			dest:  3,
			res:   25,
		},
		{
			name:  "36",
			dist:  []int{6, 47, 48, 31, 10, 27, 46, 33, 52, 33, 38, 25, 32, 45, 36, 3, 0, 33, 22, 53, 8, 13, 18, 1, 44, 41, 14, 5, 38, 25, 48},
			start: 22,
			dest:  0,
			res:   234,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, distanceBetweenBusStops(c.dist, c.start, c.dest))
		})
	}
}

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	var right int
	for i := start; ; i++ {
		if i >= len(distance) {
			i = 0
		}

		if i == destination {
			break
		}

		dist := distance[i]
		right += dist
	}

	const gap = 1
	var left int
	for i := start - 1; i != destination-gap; i-- {
		if i < 0 {
			i = len(distance) - 1
		}

		dist := distance[i]
		left += dist

		if i == destination-gap {
			break
		}
	}

	return min(right, left)
}
