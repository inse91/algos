package _1394_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-lucky-integer-in-an-array/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  int
	}{
		{
			name: "1",
			arr:  []int{2, 2, 3, 4},
			res:  2,
		}, {
			name: "2",
			arr:  []int{1, 2, 2, 3, 3, 3},
			res:  3,
		}, {
			name: "3",
			arr:  []int{2, 2, 2, 3, 3},
			res:  -1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findLucky(c.arr))
		})
	}
}

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, n := range arr {
		m[n]++
	}

	largest := -1
	for k, v := range m {
		if k == v {
			largest = max(largest, v)
		}
	}

	return largest
}
