package _747_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLargestNumberAtLeastTwiceOfOthers(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "1",
			nums: []int{3, 6, 1, 0},
			res:  1,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4},
			res:  -1,
		},
		{
			name: "3",
			nums: []int{0, 0, 3, 2},
			res:  -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, dominantIndex(c.nums))
		})
	}
}

func dominantIndex(nums []int) int {
	mx := -1
	mxIdx := -1
	prMx := -1
	for i, num := range nums {
		if num > mx {
			prMx = mx
			mx = num
			mxIdx = i
			continue
		}
		if num > prMx {
			prMx = num
		}
	}

	if prMx*2 > mx {
		return -1
	}
	return mxIdx
}
