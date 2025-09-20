package _2848_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/points-that-intersect-with-cars

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums [][]int
		res  int
	}{
		{
			name: "1",
			nums: [][]int{
				{3, 6},
				{1, 5},
				{4, 7},
			},
			res: 7,
		},
		{
			name: "2",
			nums: [][]int{
				{1, 3},
				{5, 8},
			},
			res: 7,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfPoints(c.nums))
		})
	}
}

func numberOfPoints(nums [][]int) int {
	var pts [101]int

	for _, line := range nums {
		for i := line[0]; i <= line[1]; i++ {
			pts[i]++
		}
	}

	var res int
	for i := 1; i < len(pts); i++ {
		if pts[i] > 0 {
			res++
		}
	}

	return res
}
