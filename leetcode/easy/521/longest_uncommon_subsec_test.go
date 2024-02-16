package _521_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLongestUncommonSubseq(t *testing.T) {
	cases := []struct {
		name string
		a    string
		b    string
		res  int
	}{
		{
			name: "1.1",
			a:    "aba",
			b:    "cdc",
			res:  3,
		},
		{
			name: "1.2",
			a:    "aaa",
			b:    "bbb",
			res:  3,
		},
		{
			name: "2",
			a:    "aaa",
			b:    "aaa",
			res:  -1,
		},
		{
			name: "3",
			a:    "aaa",
			b:    "a",
			res:  1,
		},
		{
			name: "4",
			a:    "a",
			b:    "a",
			res:  -1,
		},
		{
			name: "5",
			a:    "a",
			b:    "b",
			res:  1,
		},
		{
			name: "6",
			a:    "abcd",
			b:    "bca",
			res:  3,
		},
		{
			name: "7",
			a:    "aefawfawfawfaw",
			b:    "aefawfeawfwafwaef",
			res:  17,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			res := findLUSlength(c.a, c.b)
			assert.Equal(t, c.res, res)
		})
	}
}

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	if len(a) < len(b) {
		a, b = b, a
	}

	for l := len(a); l > 0; l-- {
		for i := 0; i < len(a)-l+1; i++ {
			st := i
			end := st + l
			part := a[st:end]
			if !strings.Contains(b, part) {
				return l
			}
		}
	}
	return -1
}
