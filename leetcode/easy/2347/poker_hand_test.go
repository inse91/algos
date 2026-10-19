package _2347_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/best-poker-hand

func Test(t *testing.T) {
	cases := []struct {
		name  string
		ranks []int
		suits []byte
		res   string
	}{
		{
			name:  "1",
			ranks: []int{13, 2, 3, 1, 9},
			suits: []byte{'a', 'a', 'a', 'a', 'a'},
			res:   "Flush",
		},
		{
			name:  "2",
			ranks: []int{4, 4, 2, 4, 4},
			suits: []byte{'d', 'a', 'a', 'b', 'c'},
			res:   "Three of a Kind",
		},
		{
			name:  "3",
			ranks: []int{10, 10, 2, 12, 9},
			suits: []byte{'a', 'b', 'c', 'a', 'd'},
			res:   "Pair",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, bestHand(c.ranks, c.suits))
		})
	}
}

func bestHand(ranks []int, suits []byte) string {
	s := make(map[byte]int)
	r := make(map[int]int, len(ranks))
	var (
		hasTrio bool
		hasPair bool
	)
	for i, suit := range suits {
		s[suit]++
		r[ranks[i]]++
		if r[ranks[i]] == 2 {
			hasPair = true
		}
		if r[ranks[i]] == 3 {
			hasTrio = true
		}

	}

	switch {
	case len(s) == 1:
		return "Flush"
	case hasTrio:
		return "Three of a Kind"
	case hasPair:
		return "Pair"
	default:
		return "High Card"
	}
}
