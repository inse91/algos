package _1945_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/sum-of-digits-of-string-after-convert/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		k    int
		res  int
	}{
		{
			name: "1",
			s:    "iiii",
			k:    1,
			res:  36,
		},
		{
			name: "2",
			s:    "leetcode",
			k:    2,
			res:  6,
		},
		{
			name: "3",
			s:    "zbax",
			k:    2,
			res:  8,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, getLucky(c.s, c.k))
		})
	}
}

func getLucky(s string, k int) int {
	var sb strings.Builder
	for _, r := range s {
		sb.WriteString(strconv.Itoa(int(r - 'a')))
	}

	s = sb.String()
	var acc int
	for i := range k {
		acc = 0
		for _, r := range s {
			acc += int(r - '0')
		}

		if i == k-1 {
			return acc
		}
		s = strconv.Itoa(acc)
	}

	return acc
}
