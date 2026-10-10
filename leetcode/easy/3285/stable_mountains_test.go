package _3285_test

// https://leetcode.com/problems/find-indices-of-stable-mountains

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	cases := []struct {
		name   string
		height []int
		thres  int
		res    []int
	}{
		{
			name:   "1",
			height: []int{1, 2, 3, 4, 5},
			thres:  2,
			res:    []int{3, 4},
		},
		{
			name:   "2",
			height: []int{10, 1, 10, 1, 10},
			thres:  3,
			res:    []int{1, 3},
		},
		{
			name:   "3",
			height: []int{10, 1, 10, 1, 10},
			thres:  10,
			res:    []int{},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, stableMountains(c.height, c.thres))
		})
	}
}

func stableMountains(height []int, threshold int) []int {
	var res []int
	for i := 1; i < len(height); i++ {
		prev := height[i-1]
		if prev > threshold {
			res = append(res, i)
		}
	}

	return res
}
