package _1619_test

import (
	"math"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/mean-of-array-after-removing-some-elements/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  float64
	}{
		{
			name: "1",
			arr:  []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3},
			res:  2.00000,
		},
		{
			name: "2",
			arr:  []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0},
			res:  4.00000,
		},
		{
			name: "3",
			arr:  []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4},
			res:  4.77778,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, trimMean(c.arr))
		})
	}
}

func trimMean(arr []int) float64 {
	slices.Sort(arr)

	trimLen := len(arr) / 20
	newLen := len(arr) - trimLen

	arr = arr[trimLen:newLen]

	var acc int
	for _, num := range arr {
		acc += num
	}

	return (math.Round(float64(acc) / float64(newLen-trimLen) * 100000)) / 100000
}
