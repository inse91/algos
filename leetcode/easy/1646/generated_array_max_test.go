package _1646_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/get-maximum-in-generated-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "0_0",
			n:    0,
			res:  0,
		},
		{
			name: "0_1",
			n:    1,
			res:  1,
		},
		{
			name: "1",
			n:    7,
			res:  3,
		},
		{
			name: "2",
			n:    2,
			res:  1,
		},
		{
			name: "3",
			n:    3,
			res:  2,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, getMaximumGenerated(c.n))
		})
	}
}

func getMaximumGenerated(n int) int {
	slc := []int{0, 1}
	if n <= 1 {
		return slc[n]
	}

	slc = slices.Grow(slc, n-1)
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			slc = append(slc, slc[i/2])
			continue
		}

		slc = append(slc, slc[i/2]+slc[i/2+1])
	}

	return slices.Max(slc)
}
