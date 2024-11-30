package _1582_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/special-positions-in-a-binary-matrix/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtx  [][]int
		res  int
	}{
		{
			name: "1",
			mtx:  [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}},
			res:  1,
		},
		{
			name: "1_1",
			mtx:  [][]int{{1, 1, 0}, {0, 1, 1}, {1, 1, 0}},
			res:  0,
		},
		{
			name: "3",
			mtx:  [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			res:  3,
		},
		{
			name: "66",
			mtx: [][]int{
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, // 1
				{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0}, // 2
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, // 3
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0}}, // 4
			res: 4,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numSpecial(c.mtx))
		})
	}
}

func numSpecial(mat [][]int) int {
	rowOnes := make(map[int]int)
	colOnes := make(map[int]int)

	for i, row := range mat {
		var needToBreak bool
		for j, num := range row {
			if num == 0 {
				continue
			}

			colOnes[j]++
			if _, ok := rowOnes[i]; ok {
				delete(rowOnes, i)
				needToBreak = true
			}

			if !needToBreak {
				rowOnes[i] = j
			}
		}
	}

	if len(rowOnes) == 0 {
		return 0
	}

	var res int
	for _, oneIdx := range rowOnes {
		o := colOnes[oneIdx]
		if o == 1 {
			res++
		}
	}

	return res
}
