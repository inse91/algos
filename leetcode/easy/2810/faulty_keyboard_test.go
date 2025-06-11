package _2810_test

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://leetcode.com/problems/faulty-keyboard

func Test(t *testing.T) {
	cases := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "1",
			s:    "string",
			res:  "rtsng",
		},
		{
			name: "2",
			s:    "poiinter",
			res:  "ponter",
		},
		{
			name: "3",
			s:    "goci",
			res:  "cog",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assert.Equal(t, c.res, finalString(c.s))
		})
	}
}

func finalString(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	for i := 0; i < len(s); i++ {
		bt := s[i]
		if bt != 'i' {
			sb.WriteByte(bt)
			continue
		}

		if i != len(s)-1 && s[i+1] == 'i' {
			i++
			continue
		}

		rev := reverse(sb.String())
		sb.Reset()

		sb.WriteString(rev)
	}

	return sb.String()
}

func reverse(s string) string {
	bts := []byte(s)
	slices.Reverse(bts)

	return string(bts)
}
