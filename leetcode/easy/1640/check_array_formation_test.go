package _1640_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/check-array-formation-through-concatenation/

func Test(t *testing.T) {
	cases := []struct {
		name string
		arr  []int
		pcs  [][]int
		res  bool
	}{
		{
			name: "1",
			arr:  []int{15, 88},
			pcs:  [][]int{{88}, {15}},
			res:  true,
		},
		{
			name: "2",
			arr:  []int{49, 18, 16},
			pcs:  [][]int{{16, 18, 49}},
			res:  false,
		},
		{
			name: "3",
			arr:  []int{91, 4, 64, 78},
			pcs:  [][]int{{78}, {4, 64}, {91}},
			res:  true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, canFormArray(c.arr, c.pcs))
		})
	}
}

func canFormArray(arr []int, pieces [][]int) bool {
	pcsMap := make(map[int][]int, len(pieces))
	for _, pc := range pieces {
		if len(pc) == 0 {
			continue
		}

		pcsMap[pc[0]] = pc
	}

	for i := 0; i < len(arr); {
		num := arr[i]
		pc, ok := pcsMap[num]
		if !ok {
			return false
		}

		for _, n := range pc {
			if n != arr[i] {
				return false
			}

			i++
		}
	}

	return true
}
