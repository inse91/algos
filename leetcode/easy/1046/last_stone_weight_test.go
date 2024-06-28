package _1046_test

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

// https://leetcode.com/problems/last-stone-weight/description/

func Test(t *testing.T) {
	cases := []struct {
		name   string
		stones []int
		res    int
	}{
		{
			name:   "1",
			stones: []int{2, 7, 4, 1, 8, 1},
			res:    1,
		},
		{
			name:   "2",
			stones: []int{1},
			res:    1,
		},
		{
			name:   "4",
			stones: []int{2, 2},
			res:    0,
		},
		{
			name:   "5",
			stones: []int{3, 7, 2},
			res:    2,
		},
		{
			name:   "43",
			stones: []int{5, 1, 8, 10, 7},
			res:    1,
		},
		{
			name:   "56",
			stones: []int{1, 3},
			res:    2,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, lastStoneWeight(c.stones))
		})
	}
}

func lastStoneWeight(stones []int) int {
	slices.Sort(stones)

	for {
		if len(stones) == 0 {
			return 0
		}

		if len(stones) == 1 {
			return stones[0]
		}

		newStone := stones[len(stones)-1] - stones[len(stones)-2]
		if newStone == 0 {
			stones = stones[: len(stones)-2 : len(stones)-1]
			continue
		}

		stones = rplc(stones, newStone)
	}
}

func rplc(stones []int, st int) []int {
	if len(stones) == 2 {
		return []int{st}
	}

	idx := 0
	for i := range stones {
		if stones[i] >= st {
			idx = i
			break
		}
	}

	stones = stones[: len(stones)-2 : len(stones)-1]
	idx = min(idx, len(stones))
	stones = slices.Insert(stones, idx, st)

	return stones
}
