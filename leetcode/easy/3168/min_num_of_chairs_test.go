package _3168_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-number-of-chairs-in-a-waiting-room

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "EEEEEEE",
			res:  7,
		},
		{
			name: "2",
			s:    "ELELEEL",
			res:  2,
		},
		{
			name: "3",
			s:    "ELEELEELLL",
			res:  3,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minimumChairs(c.s))
		})
	}
}

func minimumChairs(s string) int {
	var (
		res int
		cur int
	)
	for _, char := range s {
		switch char {
		case 'E':
			cur++
		case 'L':
			cur--
		default:
			panic("invalid letter: " + string(char))
		}

		res = max(res, cur)
	}

	return res
}
