package _1796_test

import (
	"testing"
	"unicode"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/second-largest-digit-in-a-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "dfa12321afd",
			res:  2,
		},
		{
			name: "1_1",
			s:    "dfa1221afd",
			res:  1,
		},
		{
			name: "2",
			s:    "abc1111",
			res:  -1,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, secondHighest(c.s))
		})
	}
}

func secondHighest(s string) int {
	var mx, pMx int32 = -1, -1
	for _, r := range s {
		if !unicode.IsDigit(r) {
			continue
		}

		dig := r - '0'
		if dig == mx {
			continue
		}

		if dig > mx {
			pMx = mx
			mx = dig

			continue
		}

		pMx = max(pMx, dig)
	}

	return int(pMx)
}
