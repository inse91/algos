package _2678_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/minimum-changes-to-make-alternating-binary-string

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "0100",
			res:  1,
		},
		{
			name: "2",
			s:    "10",
			res:  0,
		},
		{
			name: "3",
			s:    "1111",
			res:  2,
		},
		{
			name: "31",
			s:    "10010100",
			res:  3,
		},
		{
			name: "51",
			s:    "1100010111",
			res:  3,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, minOperations(c.s))
		})
	}
}

func minOperations(s string) int {
	if len(s) == 1 {
		return 0
	}
	const (
		zero = '0'
		one  = '1'
	)

	var (
		c1, c2 int
		flag   bool
	)

	for _, r := range s {
		switch r {
		case zero:
			if !flag {
				c1++
			} else {
				c2++
			}
		case one:
			if flag {
				c1++
			} else {
				c2++
			}
		}

		flag = !flag
	}

	return min(c1, c2)
}
