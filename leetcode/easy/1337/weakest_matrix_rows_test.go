package _1337_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		mtrx [][]int
		k    int
		res  []int
	}{
		{
			name: "1",
			mtrx: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:   3,
			res: []int{2, 0, 3},
		},
		{
			name: "2",
			mtrx: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			},
			k:   2,
			res: []int{0, 2},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, kWeakestRows(c.mtrx, c.k))
		})
	}
}

func kWeakestRows(mat [][]int, k int) []int {
	mapa := make(map[int][]int, len(mat))

	keys := make([]int, 0, len(mat)/4)

	for i, sols := range mat {
		var c int
		for _, sol := range sols {
			if sol == 1 {
				c++
			} else {
				break
			}
		}

		if mapa[c] == nil {
			keys = append(keys, c)
		}
		mapa[c] = append(mapa[c], i)
	}

	slices.Sort(keys)

	res := make([]int, 0, k)
outer:
	for _, key := range keys {
		rows, ok := mapa[key]
		if !ok {
			continue
		}

		for _, row := range rows {
			res = append(res, row)
			if len(res) == k {
				break outer
			}
		}
	}

	return res
}
