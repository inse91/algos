package _1103_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/distribute-candies-to-people/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		c    int
		p    int
		res  []int
	}{
		{
			name: "1",
			c:    7,
			p:    4,
			res:  []int{1, 2, 3, 1},
		},
		{
			name: "2",
			c:    10,
			p:    3,
			res:  []int{5, 2, 3},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, distributeCandies(c.c, c.p))
		})
	}
}

func distributeCandies(candies int, num_people int) []int {
	res := make([]int, num_people)

	var i int
	iter := makeSliceIter(len(res) - 1)

	for n := 1; candies > 0; n++ {
		i = iter()

		sub := min(candies, n)
		res[i] += sub

		candies -= sub
	}

	return res
}

func makeSliceIter(n int) func() int {
	i := -1

	return func() int {
		if i >= n {
			i = 0
			return i
		}

		i++
		return i
	}
}
