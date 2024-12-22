package _1624_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/largest-substring-between-two-equal-characters/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  int
	}{
		{
			name: "1",
			s:    "aa",
			res:  0,
		},
		{
			name: "2",
			s:    "abca",
			res:  2,
		},
		{
			name: "3",
			s:    "cbzxy",
			res:  -1,
		},
		{
			name: "12",
			s:    "mgntdygtxrvxjnwksqhxuxtrv",
			res:  18,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, maxLengthBetweenEqualCharacters(c.s))
		})
	}
}

func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune]int, len(s))
	res := -1
	for i, r := range s {
		v, ok := m[r]
		if !ok {
			m[r] = i
			continue
		}

		res = max(res, i-v-1)
	}

	return res
}
