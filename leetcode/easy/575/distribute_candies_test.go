package _575_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDistributeCandies(t *testing.T) {
	cases := []struct {
		name string
		c    []int
		res  int
	}{
		{
			name: "1",
			c:    []int{1, 1, 2, 3},
			res:  2,
		},
		{
			name: "2",
			c:    []int{1, 1, 2, 3, 3, 2},
			res:  3,
		},
		{
			name: "3",
			c:    []int{6, 6, 6, 6},
			res:  1,
		},
		{
			name: "4",
			c:    []int{1000, 1, 1, 1},
			res:  2,
		},
		{
			name: "5",
			c:    []int{0, 0, 14, 0, 10, 0, 0, 0},
			res:  3,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := distributeCandies(c.c)
			assert.Equal(t, c.res, res)
		})
	}
}

func distributeCandies(candyType []int) (res int) {
	set := make(map[int]struct{})

	for _, v := range candyType {
		if res == len(candyType)/2 {
			return res
		}
		if _, ok := set[v]; ok {
			continue
		}
		res++
		set[v] = struct{}{}
	}

	return res
}
