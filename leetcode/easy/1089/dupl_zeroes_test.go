package _1089_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/duplicate-zeros/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		res  []int
	}{
		{
			name: "1",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			res:  []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "2",
			arr:  []int{1, 2, 3},
			res:  []int{1, 2, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			arr := slices.Clone(c.arr)
			duplicateZeros(arr)
			assert.Equal(t, c.res, arr)
		})
	}
}

func duplicateZeros(arr []int) {
	res := make([]int, 0, len(arr))

	var containsZeroes bool

	for _, num := range arr {
		if len(res) >= len(arr) {
			break
		}

		res = append(res, num)
		if num == 0 {
			containsZeroes = true
			res = append(res, 0)
		}
	}

	if !containsZeroes {
		return
	}

	for i := range arr {
		arr[i] = res[i]
	}
}
