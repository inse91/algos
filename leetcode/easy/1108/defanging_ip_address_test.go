package _1108_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string/description/

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "1.1.1.1",
			res:  "1[.]1[.]1[.]1",
		},
		{
			name: "2",
			s:    "255.100.50.0",
			res:  "255[.]100[.]50[.]0",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, defangIPaddr(c.s))
		})
	}
}

func defangIPaddr(address string) string {
	var sb strings.Builder

	for _, r := range address {
		if r == '.' {
			sb.WriteString("[.]")
			continue
		}

		sb.WriteRune(r)
	}

	return sb.String()
}
