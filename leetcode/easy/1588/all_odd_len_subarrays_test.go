package _1588_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-all-odd-length-subarrays/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  int
	}{
		{
			name: "1",
			arr:  []int{1, 4, 2, 5, 3},
			res:  58,
		},
		{
			name: "2",
			arr:  []int{1, 2},
			res:  3,
		},
		{
			name: "2_1",
			arr:  []int{1, 2, 3},
			res:  12,
		},
		{
			name: "3",
			arr:  []int{10, 11, 12},
			res:  66,
		},
		{
			name: "4",
			arr:  []int{1, 2, 4, 5},
			res:  30,
		},
		{
			name: "5",
			arr:  []int{1, 2, 3, 4, 5, 6},
			res:  98,
		},
		{
			name: "5",
			arr:  []int{1, 2, 3, 4, 5, 6, 7},
			res:  176,
		},
		{
			name: "6",
			arr:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			res:  425,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, sumOddLengthSubarrays(c.arr))
		})
	}
}

func sumOddLengthSubarrays(arr []int) int {
	_len := len(arr)
	var inc int
	var sum int
	for i := range arr {
		for subLen := 1; i+subLen <= _len; subLen += 2 {
			for j := i; j < i+subLen; j++ {
				inc = arr[j]
				sum += inc
			}
		}
	}

	return sum
}
