package _2363_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/take-gifts-from-the-richest-pile

func Test(t *testing.T) {
	cases := []struct {
		name   string
		items1 [][]int
		items2 [][]int
		res    [][]int
	}{
		{
			name: "1",
			items1: [][]int{
				{1, 1},
				{4, 5},
				{3, 8},
			},
			items2: [][]int{
				{3, 1},
				{1, 5},
			},
			res: [][]int{
				{1, 6},
				{3, 9},
				{4, 5},
			},
		},
		{
			name: "2",
			items1: [][]int{
				{1, 1},
				{3, 2},
				{2, 3},
			},
			items2: [][]int{
				{2, 1},
				{3, 2},
				{1, 3},
			},
			res: [][]int{
				{1, 4},
				{2, 4},
				{3, 4},
			},
		},
		{
			name: "3",
			items1: [][]int{
				{1, 3},
				{2, 2},
			},
			items2: [][]int{
				{7, 1},
				{2, 2},
				{1, 4},
			},
			res: [][]int{
				{1, 7},
				{2, 4},
				{7, 1},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, mergeSimilarItems(c.items1, c.items2))
		})
	}
}

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	m := make(map[int]int, max(len(items1), len(items2)))
	values := make([]int, 0, max(len(items1), len(items2)))
	for _, item := range append(items1, items2...) {
		value, weight := item[0], item[1]
		if _, ok := m[value]; !ok {
			values = append(values, value)
		}
		m[value] += weight
	}

	slices.Sort(values)
	res := make([][]int, 0, len(m))
	for _, value := range values {
		res = append(res, []int{value, m[value]})
	}

	return res
}
