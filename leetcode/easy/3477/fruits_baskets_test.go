package _3477_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/fruits-into-baskets-ii/description

func Test(t *testing.T) {
	cases := []struct {
		name    string
		fruits  []int
		baskets []int
		res     int
	}{
		{
			name:    "1",
			fruits:  []int{4, 2, 5},
			baskets: []int{3, 5, 4},
			res:     1,
		},
		{
			name:    "2",
			fruits:  []int{3, 6, 1},
			baskets: []int{6, 4, 7},
			res:     0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numOfUnplacedFruits(c.fruits, c.baskets))
		})
	}
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	if len(fruits) == 0 || len(baskets) == 0 {
		return 0
	}

	takenBasketsIndexes := make(map[int]struct{}, len(baskets))
	for _, f := range fruits {
		for j, b := range baskets {
			if f > b {
				continue
			}

			if _, ok := takenBasketsIndexes[j]; ok {
				continue
			}

			takenBasketsIndexes[j] = struct{}{}
			break
		}
	}

	return len(baskets) - len(takenBasketsIndexes)
}
