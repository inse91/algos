package _2138_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/divide-a-string-into-groups-of-size-k

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		k    int
		fill byte
		res  []string
	}{
		{
			name: "1",
			s:    "abcdefghi",
			k:    3,
			fill: 'x',
			res:  []string{"abc", "def", "ghi"},
		},
		{
			name: "2",
			s:    "abcdefghij",
			k:    3,
			fill: 'x',
			res:  []string{"abc", "def", "ghi", "jxx"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.ElementsMatch(t, c.res, divideString(c.s, c.k, c.fill))
		})
	}
}

func divideString(s string, k int, fill byte) []string {
	res := make([]string, 0, len(s)/k)
	var (
		c  int
		sb strings.Builder
	)

	for _, r := range s {
		sb.WriteRune(r)
		c++
		if c != k {
			continue
		}

		c = 0
		res = append(res, sb.String())
		sb.Reset()
	}

	if len(s)%k == 0 {
		return res
	}

	for range k - c {
		sb.WriteByte(fill)
	}
	res = append(res, sb.String())

	return res
}
