package _1385_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/find-the-distance-value-between-two-arrays/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr1 []int
		arr2 []int
		d    int
		res  int
	}{
		{
			name: "1",
			arr1: []int{4, 5, 8},
			arr2: []int{10, 9, 1, 8},
			d:    2,
			res:  2,
		},
		{
			name: "2",
			arr1: []int{1, 4, 2, 3},
			arr2: []int{-4, -3, 6, 10, 20, 30},
			d:    3,
			res:  2,
		},
		{
			name: "3",
			arr1: []int{2, 1, 100, 3},
			arr2: []int{-5, -2, 10, -3, 7},
			d:    6,
			res:  1,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, findTheDistanceValue(c.arr1, c.arr2, c.d))
		})
	}
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	set := make(map[int]struct{}, len(arr2))
	for _, i := range arr2 {
		set[i] = struct{}{}
	}

	var res int
out:
	for _, num := range arr1 {
		var ok bool
		if _, ok = set[num]; ok {
			continue out
		}
		for i := 1; i <= d; i++ {
			numPlus := num + i
			if _, ok = set[numPlus]; ok {
				continue out
			}
			numMinus := num - i
			if _, ok = set[numMinus]; ok {
				continue out
			}
		}

		res++
	}

	return res
}
