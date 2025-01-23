package _1688_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-of-matches-in-tournament/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		n    int
		res  int
	}{
		{
			name: "1",
			n:    7,
			res:  6,
		},
		{
			name: "2",
			n:    14,
			res:  13,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numberOfMatches(c.n))
		})
	}
}

func numberOfMatches(n int) int {
	var c int
	for {
		if n == 1 {
			break
		}

		if n == 2 {
			c++
			break
		}

		if n%2 == 0 {
			n /= 2
			c += n
			continue
		}

		n = (n-1)/2 + 1
		c += n - 1
	}

	return c
}
