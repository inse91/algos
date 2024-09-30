package _1356_test

import (
	"math/bits"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sort-integers-by-the-number-of-1-bits/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		slc  []int
		res  []int
	}{
		{
			name: "1",
			slc:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
			res:  []int{0, 1, 2, 4, 8, 3, 5, 6, 7},
		}, {
			name: "2",
			slc:  []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
			res:  []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, sortByBits(c.slc))
			for i, num := range sortByBits(c.slc) {
				assert.Equal(t, c.res[i], num)
			}
		})
	}
}

func sortByBits(arr []int) []int {
	slices.SortFunc(arr, func(a, b int) int {
		dif := bits.OnesCount32(uint32(a)) - bits.OnesCount32(uint32(b))
		if dif == 0 {
			return a - b
		}

		return dif
	})

	return arr
}
