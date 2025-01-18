package _1560_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/most-visited-sector-in-a-circular-track/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		n      int
		rounds []int
		res    []int
	}{
		{
			name:   "1",
			n:      4,
			rounds: []int{1, 3, 1, 2},
			res:    []int{1, 2},
		},
		{
			name:   "1_",
			n:      4,
			rounds: []int{1, 3, 1, 3},
			res:    []int{1, 2, 3},
		},
		{
			name:   "2",
			n:      4,
			rounds: []int{2, 1, 2, 1, 2, 1, 2, 1, 2},
			res:    []int{2},
		},
		{
			name:   "3",
			n:      7,
			rounds: []int{1, 3, 5, 7},
			res:    []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:   "4",
			n:      4,
			rounds: []int{3, 1, 2},
			res:    []int{1, 2, 3, 4},
		},
		{
			name:   "4_",
			n:      4,
			rounds: []int{3, 1, 4},
			res:    []int{3, 4},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, mostVisited(c.n, c.rounds))
		})
	}
}

func mostVisited(n int, rounds []int) []int {
	start := rounds[0]
	end := rounds[len(rounds)-1]
	res := []int{end}
	if start == end {
		return res
	}

	res = slices.Grow(res, len(res)-1)
	it := iter{
		size: n,
		cur:  rounds[len(rounds)-1],
	}

	for {
		res = append(res, it.prev())
		if res[len(res)-1] == start {
			break
		}
	}

	slices.Sort(res)
	return res
}

type iter struct {
	size int
	cur  int
}

func (i *iter) prev() int {
	i.cur--
	if i.cur == 0 {
		i.cur = i.size
	}

	return i.cur
}
