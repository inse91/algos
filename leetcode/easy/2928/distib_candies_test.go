package _2928_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/distribute-candies-among-children-i

func Test(t *testing.T) {
	cases := []struct {
		name  string
		n     int
		limit int
		res   int
	}{
		{
			name:  "1",
			n:     5,
			limit: 2,
			res:   3,
		},
		{
			name:  "2",
			n:     3,
			limit: 3,
			res:   10,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, distributeCandies(c.n, c.limit))
		})
	}
}

func distributeCandies(n int, limit int) int {
	set := make(map[[3]int]struct{}, n*limit)
	for i := 0; i <= limit; i++ {
		for j := 0; j <= limit; j++ {
			for k := 0; k <= limit; k++ {
				summ := i + j + k
				if summ > n {
					break
				}
				if summ == n {
					set[[3]int{i, j, k}] = struct{}{}
				}
			}
		}
	}

	return len(set)
}
