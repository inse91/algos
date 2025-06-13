package _2206_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/divide-array-into-equal-pairs

func Test(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		res  bool
	}{
		{
			name: "1",
			nums: []int{3, 2, 3, 2, 2, 2},
			res:  true,
		},
		{
			name: "2",
			nums: []int{1, 2, 3, 4},
			res:  false,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, divideArray(c.nums))
		})
	}
}

func divideArray(nums []int) bool {
	m := make(map[int]int, len(nums)/2)
	for _, num := range nums {
		m[num]++
	}

	for _, q := range m {
		if q%2 != 0 {
			return false
		}
	}

	return true
}
