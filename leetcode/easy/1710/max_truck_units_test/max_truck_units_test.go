package _1710_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/maximum-units-on-a-truck/description/

func Test(t *testing.T) {
	cases := []struct {
		name  string
		boxes [][]int
		size  int
		res   int
	}{
		{
			name: "1",
			boxes: [][]int{
				{1, 3},
				{2, 2},
				{3, 1},
			},
			size: 4,
			res:  8,
		},
		{
			name: "2",
			boxes: [][]int{
				{5, 10},
				{2, 5},
				{4, 7},
				{3, 9},
			},
			size: 10,
			res:  91,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maximumUnits(c.boxes, c.size))
		})
	}
}

func maximumUnits(boxTypes [][]int, truckSize int) int {
	slices.SortFunc(boxTypes, func(a, b []int) int {
		if min(len(a), len(b)) < 2 {
			panic("invalid input")
		}
		if a[1] < b[1] {
			return 1
		}

		return -1
	})

	var acc int
	for _, boxInfo := range boxTypes {
		boxCount := boxInfo[0]
		boxSize := boxInfo[1]

		if truckSize <= boxCount {
			actualCount := min(boxCount, truckSize)
			acc += actualCount * boxSize

			break
		}

		truckSize -= boxCount
		acc += boxCount * boxSize
	}

	return acc
}
