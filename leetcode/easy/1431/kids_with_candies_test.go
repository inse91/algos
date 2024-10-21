package _1431_test

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/description/

func Test(t *testing.T) {
	cases := []struct {
		name    string
		candies []int
		extra   int
		res     []bool
	}{
		{
			name:    "1",
			candies: []int{2, 3, 5, 1, 3},
			extra:   3,
			res:     []bool{true, true, true, false, true},
		},
		{
			name:    "2",
			candies: []int{4, 2, 1, 1, 2},
			extra:   1,
			res:     []bool{true, false, false, false, false},
		},
		{
			name:    "3",
			candies: []int{12, 1, 12},
			extra:   10,
			res:     []bool{true, false, true},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, kidsWithCandies(c.candies, c.extra))
		})
	}
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxNumberOfCandies := slices.Max(candies)
	res := make([]bool, len(candies))

	for i, candy := range candies {
		res[i] = candy+extraCandies >= maxNumberOfCandies
	}

	return res
}
