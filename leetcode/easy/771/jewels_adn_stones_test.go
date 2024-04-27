package _771_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJewelsAndStones(t *testing.T) {
	cases := []struct {
		name string
		j    string
		s    string
		res  int
	}{
		{
			name: "1",
			j:    "aA",
			s:    "aAAbbbb",
			res:  3,
		},
		{
			name: "2",
			j:    "z",
			s:    "ZZ",
			res:  0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, numJewelsInStones(c.j, c.s))
		})
	}
}

func numJewelsInStones(jewels string, stones string) int {
	jSet := make(map[rune]struct{}, len(jewels))
	for _, r := range jewels {
		jSet[r] = struct{}{}
	}

	var res int
	for _, r := range stones {
		if _, ok := jSet[r]; ok {
			res++
		}
	}

	return res
}
