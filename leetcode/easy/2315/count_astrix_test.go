package _2315_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/count-asterisks/description

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "l|*e*et|c**o|*de|",
			res:  2,
		},
		{
			name: "2",
			s:    "abc",
			res:  0,
		},
		{
			name: "3",
			s:    "yo|uar|e**|b|e***au|tifu|l",
			res:  5,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, countAsterisks(c.s))
		})
	}
}

func countAsterisks(s string) int {
	var res int
	var openedBracket bool
	for _, r := range s {
		if r == '|' {
			openedBracket = !openedBracket
			continue
		}

		if openedBracket {
			continue
		}

		if r == '*' {
			res++
		}

	}

	return res
}
