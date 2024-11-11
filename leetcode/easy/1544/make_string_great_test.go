package _1544_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/make-the-string-great/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "leEeetcode",
			res:  "leetcode",
		},
		{
			name: "2",
			s:    "abBAcC",
			res:  "",
		},
		{
			name: "75_1",
			s:    "Nn",
			res:  "",
		},
		{
			name: "75_2",
			s:    "nN",
			res:  "",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, makeGood(c.s))
		})
	}
}

func makeGood(s string) string {
	const d1, d2 = 32, 224
outer:
	for {
		for i := range s {
			if i == 0 {
				continue
			}

			diff := mod(s[i], s[i-1])
			if diff == d1 || diff == d2 {
				s = clip(s, i)
				continue outer
			}
		}

		return s
	}
}

func mod(a, b byte) int {
	m := a - b
	if m < 0 {
		return int(-m)
	}

	return int(m)
}

func clip(s string, i int) string {
	if i+1 > len(s) || i-1 < 0 {
		return s
	}
	return s[:i-1] + s[i+1:]
}

func TestClip(t *testing.T) {
	s := clip("12345", 4)
	fmt.Println(s)
}
