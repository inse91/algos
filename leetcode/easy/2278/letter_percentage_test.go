package _2278_test

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/percentage-of-letter-in-string

func Test(t *testing.T) {
	cases := []struct {
		name   string
		s      string
		letter byte
		res    int
	}{
		{
			name:   "1",
			s:      "foobar",
			letter: 'o',
			res:    33,
		},
		{
			name:   "2",
			s:      "jjj",
			letter: 'k',
			res:    0,
		},
		{
			name:   "3",
			s:      "zzy",
			letter: 'z',
			res:    66,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.EqualValues(t, c.res, percentageLetter(c.s, c.letter))
		})
	}
}

func percentageLetter(s string, letter byte) int {
	var c int
	for _, r := range s {
		if byte(r) == letter {
			c++
		}
	}

	return int(math.Floor(float64(c) / float64(len(s)) * 100))
}
