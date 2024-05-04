package _908_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		k    int
		res  int
	}{
		{
			name: "1",
			nums: []int{1},
			k:    0,
			res:  0,
		},
		{
			name: "2",
			nums: []int{0, 10},
			k:    2,
			res:  6,
		},
		{
			name: "3",
			nums: []int{1, 3, 6},
			k:    3,
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, smallestRangeI(c.nums, c.k))
		})
	}
}

func smallestRangeI(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	slices.Sort(nums)

	mn := 10_000 + 1
	mx := -1

	for _, n := range nums {
		mn = min(mn, n)
		mx = max(mx, n)
	}

	mn += k
	mx -= k

	r := mx - mn
	if r < 0 {
		return 0
	}

	return r
}
