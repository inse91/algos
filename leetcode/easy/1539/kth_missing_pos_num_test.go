package _1539_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kth-missing-positive-number/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		k    int
		res  int
	}{
		{
			name: "1",
			arr:  []int{2, 3, 4, 7, 11},
			k:    5,
			res:  9,
		},
		{
			name: "1_1",
			arr:  []int{2, 3, 4, 7, 11},
			k:    6,
			res:  10,
		},
		{
			name: "1_2",
			arr:  []int{2, 3, 4, 7, 10, 11},
			k:    6,
			res:  12,
		},
		{
			name: "2",
			arr:  []int{1, 2, 3, 4},
			k:    2,
			res:  6,
		},
		{
			name: "22",
			arr:  []int{1, 3, 5},
			k:    4,
			res:  7,
		},
		{
			name: "999",
			arr:  []int{30},
			k:    25,
			res:  25,
		},
		{
			name: "999_",
			arr:  []int{1},
			k:    25,
			res:  26,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findKthPositive(c.arr, c.k))
		})
	}
}

func findKthPositive(arr []int, k int) int {
	_len := len(arr)
	lastOne := arr[_len-1]
	missingOnes := lastOne - _len

	if missingOnes < k {
		return lastOne + k - missingOnes
	}

	var c int
	cur := 1
	for i := 0; i < _len; {
		_i := arr[i]
		if _i == cur {
			cur++
			i++
			continue
		}

		c++
		if c == k {
			return cur
		}
		cur++
	}

	return 0
}
